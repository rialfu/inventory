package service

import (
	"context"
	"fmt"
	"rialfu/wallet/database/entities"
	"rialfu/wallet/modules/category/dto"
	"rialfu/wallet/modules/category/repository"
	"rialfu/wallet/pkg/constants"
	"rialfu/wallet/pkg/helpers"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type CategoryService interface {
	Create(ctx context.Context, ins dto.CategoryCreateRequest) (dto.CategoryResponse, error)
	GetDropdown(ctx context.Context, search string, page int) ([]dto.CategoryDropdownDTO, error)
	Update(ctx context.Context, data dto.CategoryUpdateRequest, id string) (dto.CategoryResponse, error)
	GetAll(ctx context.Context, id *string, queryParam map[string][]string) (helpers.PaginateData[dto.CategoryResponse], *string, error)
	Get(ctx context.Context, id string) (dto.CategoryResponse, error)
}

type categoryService struct {
	repository repository.CategoryRepository
	db         *gorm.DB
	transactor helpers.Transactor
}

func NewCategoryService(
	catRepo repository.CategoryRepository,
	db *gorm.DB,
) CategoryService {
	transactor := helpers.NewGormTransactor(db)
	return &categoryService{
		repository: catRepo,
		db:         db,
		transactor: transactor,
	}
}
func (s *categoryService) Create(ctx context.Context, ins dto.CategoryCreateRequest) (dto.CategoryResponse, error) {
	data := entities.Category{
		Name:  ins.Name,
		Depth: 0,
	}
	var parentID uint64
	var err error
	if ins.ParentID != "" {
		parentID, err = strconv.ParseUint(ins.ParentID, 10, 64)
		if err != nil {
			return dto.CategoryResponse{}, dto.ErrGetParentNotFound
		}
	}

	res, err := s.repository.Create(ctx, data, parentID)
	if err != nil {
		return dto.CategoryResponse{}, err
	}
	display := s.displayPath(res.Path)
	return dto.CategoryResponse{ID: res.ID, Name: res.Name, ParentID: res.ParentID, Path: &display}, nil
}
func (s *categoryService) GetDropdown(ctx context.Context, search string, page int) ([]dto.CategoryDropdownDTO, error) {
	fmt.Println("search", search)
	datas, err := s.repository.ReadDropdown(ctx, search, 10, page)
	if err != nil {
		return []dto.CategoryDropdownDTO{}, err
	}
	res := make([]dto.CategoryDropdownDTO, len(datas))
	for i, data := range datas {
		res[i] = dto.CategoryDropdownDTO{ID: data.ID, Name: s.displayPath(data.Path)}
	}
	return res, nil
}

func (s *categoryService) GetAll(ctx context.Context, id *string, queryParam map[string][]string) (helpers.PaginateData[dto.CategoryResponse], *string, error) {
	var parentPath *string
	datas, page, limit, total, err := s.repository.ReadAll(ctx, queryParam, id)

	if err != nil {
		return helpers.PaginateData[dto.CategoryResponse]{}, parentPath, nil
	}
	results := make([]dto.CategoryResponse, 0, len(datas))
	for _, u := range datas {
		item := dto.CategoryResponse{
			ID:       u.ID,
			Name:     u.Name,
			Path:     &u.Path,
			ParentID: u.ParentID,
		}
		results = append(results, item)
	}
	if id != nil && len(datas) > 0 {
		parts := strings.Split(datas[0].Path, "/")
		if len(parts) > 1 {
			joinedPath := strings.Join(parts[:len(parts)-1], "/")
			parentPath = &joinedPath
		}
	} else if id != nil {
		parent, exist, err := s.repository.GetById(ctx, *id, false)
		if err != nil {
			return helpers.PaginateData[dto.CategoryResponse]{}, parentPath, err
		}
		if exist == false {
			return helpers.PaginateData[dto.CategoryResponse]{Data: results, Page: page, Limit: limit, Total: total}, parentPath, nil
		}
		parentPath = &parent.Path

	}
	return helpers.PaginateData[dto.CategoryResponse]{Data: results, Page: page, Limit: limit, Total: total}, parentPath, nil
}
func (s *categoryService) Get(ctx context.Context, id string) (dto.CategoryResponse, error) {
	var res dto.CategoryResponse
	data, exist, err := s.repository.GetById(ctx, id, false)
	if err != nil {
		return res, err
	}
	if exist == false {
		return res, constants.ErrDataNotFound
	}
	res.ID = data.ID
	res.Name = data.Name
	res.Path = &data.Path
	return res, nil
}
func (s *categoryService) Update(ctx context.Context, data dto.CategoryUpdateRequest, id string) (dto.CategoryResponse, error) {
	var res dto.CategoryResponse
	var newParentID *uint64
	if data.ParentID != "" && data.ParentID != "0" {
		val, err := strconv.ParseUint(data.ParentID, 10, 64)
		if err != nil {
			return res, dto.ErrGetParentNotFound
		}
		newParentID = &val
	}
	var sb strings.Builder
	if data.ParentID == id {
		return res, dto.ErrDataCircular
	}

	err := s.transactor.WithTransaction(ctx, func(txCtx context.Context) error {
		dataInDB, isExist, err := s.repository.GetById(txCtx, id, true)

		if err != nil {
			return err
		} else if isExist == false {
			return constants.ErrDataNotFound
		}

		oldPath := dataInDB.Path
		oldDepth := int(dataInDB.Depth)

		var newPath string
		var newDepth int

		if newParentID != nil {
			parent, isExist, err := s.repository.GetById(ctx, data.ParentID, true)
			if err != nil {
				return err
			}
			if isExist == false {
				return dto.ErrGetParentNotFound
			}
			if strings.HasPrefix(parent.Path, dataInDB.Path) {
				return dto.ErrDataCircular
			}
			newDepth = int(parent.Depth) + 1

			sb.WriteString(parent.Path)
			sb.WriteString("/")
			sb.WriteString(strconv.FormatUint(dataInDB.ID, 10))
			sb.WriteString("-")
			sb.WriteString(data.Name)
			newPath = sb.String()

			// newPath = parent.Path + "/" + padID(cat.ID) + "-" + newSlug
		} else {
			newDepth = 0
			sb.WriteString(strconv.FormatUint(dataInDB.ID, 10))
			sb.WriteString("-")
			sb.WriteString(data.Name)
			newPath = sb.String()
		}
		sb.Reset()

		err = s.repository.Update(txCtx, id, map[string]any{"name": data.Name, "path": newPath, "depth": newDepth, "parent_id": newParentID})
		if err != nil {
			return err
		}
		err = s.repository.UpdateSubtreePath(txCtx, oldPath, newPath)
		if err != nil {
			return err
		}
		err = s.repository.UpdateSubtreeDepth(txCtx, oldDepth, newDepth, newPath)
		if err != nil {
			return err
		}
		val, err := strconv.ParseUint(data.ParentID, 10, 64)
		display := s.displayPath(newPath)
		res = dto.CategoryResponse{
			ID:       val,
			Name:     data.Name,
			Path:     &display,
			ParentID: newParentID,
		}
		return nil
	})
	return res, err
}

func (s *categoryService) displayPath(path string) string {
	parts := strings.Split(path, "/")
	for i, p := range parts {
		if idx := strings.Index(p, "-"); idx != -1 {
			parts[i] = p[idx+1:]
		}
	}
	return strings.Join(parts, "/")
}
