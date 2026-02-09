package service

import (
	"context"
	"rialfu/wallet/database/entities"
	"rialfu/wallet/modules/merk/dto"
	"rialfu/wallet/modules/merk/repository"
	"rialfu/wallet/pkg/constants"
	"rialfu/wallet/pkg/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MerkService interface {
	Create(ctx context.Context, req dto.MerkCreateRequest) (dto.MerkResponse, error)
	Update(ctx context.Context, req dto.MerkUpdateRequest, id string) (dto.MerkResponse, error)
	GetAll(ctx *gin.Context) (helpers.PaginateData[dto.MerkResponse], error)
	Get(ctx *gin.Context, id string) (dto.MerkResponse, error)
}

type merkService struct {
	repository repository.MerkRepository
	db         *gorm.DB
}

func NewMerkService(
	repo repository.MerkRepository,
	db *gorm.DB,
) MerkService {
	return &merkService{
		repository: repo,
		db:         db,
	}
}

func (s *merkService) Create(ctx context.Context, req dto.MerkCreateRequest) (dto.MerkResponse, error) {
	_, isExist, err := s.repository.GetByName(ctx, req.Name)
	if err != nil {
		return dto.MerkResponse{}, err
	}
	if isExist {
		return dto.MerkResponse{}, constants.ErrValueNotUniq
	}
	data := entities.Merk{
		Name: req.Name,
	}
	updated, err := s.repository.Create(ctx, data)
	res := dto.MerkResponse{
		ID:   updated.ID,
		Name: updated.Name,
	}

	return res, err
}

func (s *merkService) GetAll(ctx *gin.Context) (helpers.PaginateData[dto.MerkResponse], error) {
	datas, page, limit, total, err := s.repository.ReadAll(ctx, ctx.Request.URL.Query())

	if err != nil {
		return helpers.PaginateData[dto.MerkResponse]{}, err
	}
	results := make([]dto.MerkResponse, 0, len(datas))
	for _, u := range datas {
		item := dto.MerkResponse{
			ID:   u.ID,
			Name: u.Name,
		}
		results = append(results, item)
	}

	return helpers.PaginateData[dto.MerkResponse]{Data: results, Limit: limit, Page: page, Total: total}, nil
}

func (s *merkService) Get(ctx *gin.Context, id string) (dto.MerkResponse, error) {
	var err error
	var res dto.MerkResponse

	data, isExist, err := s.repository.GetById(ctx, id)
	if err != nil {
		return res, err
	}
	if isExist == false {
		return res, constants.ErrDataNotFound
	}
	res.ID = data.ID
	res.Name = data.Name
	return res, nil
}

func (s *merkService) Update(ctx context.Context, req dto.MerkUpdateRequest, id string) (dto.MerkResponse, error) {
	data, isExist, err := s.repository.GetByNameOrId(ctx, id, req.Name)
	if err != nil {
		return dto.MerkResponse{}, err
	}
	if isExist == false {
		if strconv.FormatUint(data.ID, 10) != id {
			return dto.MerkResponse{}, constants.ErrValueNotUniq
		}

	}

	data.Name = req.Name

	updated, err := s.repository.Update(ctx, data)
	res := dto.MerkResponse{
		ID:   updated.ID,
		Name: updated.Name,
	}

	return res, err
}
