package service

import (
	"context"
	"rialfu/wallet/database/entities"
	"rialfu/wallet/modules/item/dto"
	itemRepo "rialfu/wallet/modules/item/repository"
	"strconv"
	"strings"

	attrRepo "rialfu/wallet/modules/attribute/repository"
	dtoCat "rialfu/wallet/modules/category/dto"
	cateRepo "rialfu/wallet/modules/category/repository"
	dtoMerk "rialfu/wallet/modules/merk/dto"
	merkRepo "rialfu/wallet/modules/merk/repository"
	"rialfu/wallet/pkg/constants"
	"rialfu/wallet/pkg/helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemService interface {
	GetAll(ctx *gin.Context) (helpers.PaginateData[dto.ItemResponse], error)
	CreateOrUpItem(ctx context.Context, req dto.ItemCreateRequest) (dto.ItemResponse, error)
	// Update(ctx context.Context, req dto.MerkUpdateRequest, id string) (dto.MerkResponse, error)
	GetDropdownMerk(ctx context.Context, search string, page int) ([]dtoMerk.MerkDropdownDTO, error)
	GetDropdownCategory(ctx context.Context, search string, page int) ([]dtoCat.CategoryDropdownDTO, error)
}

type itemService struct {
	itemRepository    itemRepo.ItemRepository
	variantRepository itemRepo.ItemVariantRepository
	mr                merkRepo.MerkRepository
	cr                cateRepo.CategoryRepository
	avr               attrRepo.AttributeValueRepository
	db                *gorm.DB
	transactor        helpers.Transactor
}

func NewItemService(
	ir itemRepo.ItemRepository,
	vr itemRepo.ItemVariantRepository,
	mr merkRepo.MerkRepository,
	cr cateRepo.CategoryRepository,
	avr attrRepo.AttributeValueRepository,
	db *gorm.DB,

) ItemService {
	tr := helpers.NewGormTransactor(db)
	return &itemService{
		itemRepository:    ir,
		variantRepository: vr,
		mr:                mr,
		db:                db,
		cr:                cr,
		avr:               avr,
		transactor:        tr,
	}
}

func (s *itemService) CreateOrUpItem(ctx context.Context, req dto.ItemCreateRequest) (dto.ItemResponse, error) {

	var res dto.ItemResponse
	err := s.transactor.WithTransaction(ctx, func(txCtx context.Context) error {
		var data entities.Item
		var err error
		// mErr := &utils.MultiError{}
		isExist := false
		merk, isExist, err := s.mr.GetById(txCtx, strconv.FormatUint(req.Merk, 10))
		if err != nil {
			return err
		}
		if isExist == false {
			return dto.ErrMerkNotFound
		}

		cat, isExist, err := s.cr.GetById(txCtx, strconv.FormatUint(req.Category, 10), false)
		if err != nil {
			return err
		}
		if isExist == false {
			return dto.ErrCategoryNotFound
		}

		if req.ID == nil {
			data.Name = req.Name
			data.MerkID = merk.ID
			data.CategoryID = cat.ID
			data, err = s.itemRepository.Create(txCtx, data)
		} else {
			data, isExist, err = s.itemRepository.GetById(ctx, *req.ID, true)
			if err != nil {
				return err
			}
			if isExist == false {
				return constants.ErrDataNotFound
			}
			data.Name = req.Name
			data.MerkID = merk.ID
			data.CategoryID = cat.ID
			data, err = s.itemRepository.Update(txCtx, data)
		}
		if err != nil {
			return err
		}

		res = dto.ItemResponse{
			ID:           data.ID,
			Name:         data.Name,
			CategoryID:   data.CategoryID,
			MerkID:       data.MerkID,
			MerkName:     merk.Name,
			CategoryName: s.cleanPath(cat.Path),
		}
		return nil
	})

	return res, err
}

func (s *itemService) GetAll(ctx *gin.Context) (helpers.PaginateData[dto.ItemResponse], error) {
	var err error
	var limit, page int
	var total int64

	datas, page, limit, total, err := s.itemRepository.ReadAll(ctx, ctx.Request.URL.Query())

	if err != nil {
		return helpers.PaginateData[dto.ItemResponse]{}, err
	}
	results := make([]dto.ItemResponse, 0, len(datas))
	for _, u := range datas {
		item := dto.ItemResponse{
			ID:           u.ID,
			Name:         u.Name,
			CategoryID:   u.CategoryID,
			CategoryName: u.Category.Name,
			MerkID:       u.MerkID,
			MerkName:     u.Merk.Name,
		}
		results = append(results, item)
	}

	return helpers.PaginateData[dto.ItemResponse]{Data: results, Limit: limit, Page: page, Total: total}, err
}
func (s *itemService) CreateItemVariant(ctx context.Context, req dto.ItemVariantCreateRequest) (dto.VariantResponse, error) {

	var res dto.VariantResponse
	err := s.transactor.WithTransaction(ctx, func(txCtx context.Context) error {
		if req.Item == nil {
			return constants.ErrDataNotFound
		}
		dataItem, isExist, err := s.itemRepository.GetById(txCtx, *req.Item, true)
		if err != nil {
			return err
		}
		if isExist == false {
			return constants.ErrDataNotFound
		}
		_, isExist, err = s.variantRepository.GetByName(txCtx, req.SKU)
		if err != nil {
			return err
		}
		if isExist {
			return constants.ErrValueNotUniq
		}
		data := entities.ItemVariant{
			SKU:         req.SKU,
			ItemID:      dataItem.ID,
			Description: req.Description,
		}
		ids := make([]string, 0, len(req.Attributes))
		for _, id := range req.Attributes {
			ids = append(ids, strconv.FormatUint(id, 10))
		}
		attrs, err := s.avr.GetByIdsWithAttrName(txCtx, ids)
		if err != nil {
			return err
		}

		for _, attr := range attrs {
			data.Attr = append(data.Attr,
				entities.ItemVariantAttributeValue{AttributeValueID: attr.ID},
			)
		}
		data, err = s.variantRepository.Create(txCtx, data)
		if err != nil {
			return err
		}

		res.ID = data.ID
		res.SKU = data.SKU
		res.Description = data.Description

		groupedMap := make(map[uint64]*dto.AttributeNameResponse)
		var order []uint64

		for _, attr := range attrs {
			nameID := attr.AttributeName.ID
			if _, exist := groupedMap[nameID]; !exist {
				groupedMap[nameID] = &dto.AttributeNameResponse{
					Name:  attr.AttributeName.AttributeName, // Pastikan field ini terisi dari DB
					Value: []dto.AttributeValueResponse{},
				}
				order = append(order, nameID)
			}
			groupedMap[nameID].Value = append(groupedMap[nameID].Value, dto.AttributeValueResponse{
				Name: attr.AttributeValue,
			})
		}

		for _, id := range order {
			res.Attributes = append(res.Attributes, *groupedMap[id])
		}
		return nil
	})
	return res, err
}
func (s *itemService) UpdateItemVariant(ctx context.Context, req dto.ItemVariantCreateRequest) (dto.VariantResponse, error) {

	var res dto.VariantResponse
	err := s.transactor.WithTransaction(ctx, func(txCtx context.Context) error {
		if req.ID == nil {
			return constants.ErrDataNotFound
		}
		data, isExist, err := s.variantRepository.GetById(txCtx, strconv.FormatUint(*req.ID, 10), true)
		if err != nil {
			return err
		}
		if isExist == false {
			return constants.ErrDataNotFound
		}

		skuExist, isExist, err := s.variantRepository.GetByName(txCtx, req.SKU)
		if err != nil {
			return err
		}
		if isExist {
			if skuExist.ID != data.ID {
				return constants.ErrValueNotUniq
			}

		}
		data.SKU = req.SKU
		data.Description = req.Description

		ids := make([]string, 0, len(req.Attributes))
		for _, id := range req.Attributes {
			ids = append(ids, strconv.FormatUint(id, 10))
		}
		attrs, err := s.avr.GetByIdsWithAttrName(txCtx, ids)
		if err != nil {
			return err
		}

		data, err = s.variantRepository.Update(txCtx, data)
		if err != nil {
			return err
		}
		ivavs := make([]entities.ItemVariantAttributeValue, 0, len(attrs))
		for _, attr := range attrs {
			ivavs = append(ivavs, entities.ItemVariantAttributeValue{
				ItemVariantID:    data.ID,
				AttributeValueID: attr.ID,
			})
		}
		err = s.variantRepository.SyncVariantAttribute(txCtx, ivavs, data.ID)
		if err != nil {
			return err
		}

		res.ID = data.ID
		res.SKU = data.SKU
		res.Description = data.Description

		uniqueAttr := make(map[uint64]*dto.AttributeNameResponse) // Pakai pointer biar mudah update
		var sortOrder []uint64                                    // [TIPS] Untuk menjaga urutan atribut agar tidak acak

		for _, attr := range attrs {
			key := attr.AttributeName.ID

			// Jika belum ada di map, inisialisasi
			if _, exist := uniqueAttr[key]; !exist {
				uniqueAttr[key] = &dto.AttributeNameResponse{
					Name:  attr.AttributeName.AttributeName,
					Value: []dto.AttributeValueResponse{},
				}
				sortOrder = append(sortOrder, key)
			}

			// Append value ke slice yang ada di dalam map
			uniqueAttr[key].Value = append(uniqueAttr[key].Value, dto.AttributeValueResponse{
				Name: attr.AttributeValue, // Nama nilai (misal: "Red")
			})
		}
		res.Attributes = make([]dto.AttributeNameResponse, 0) // Reset slice
		for _, key := range sortOrder {
			res.Attributes = append(res.Attributes, *uniqueAttr[key])
		}

		return nil
	})
	return res, err
}

func (s *itemService) GetDropdownCategory(ctx context.Context, search string, page int) ([]dtoCat.CategoryDropdownDTO, error) {
	var res []dtoCat.CategoryDropdownDTO
	datas, err := s.cr.ReadDropdown(ctx, search, 10, page)
	if err != nil {
		return res, err
	}
	res = make([]dtoCat.CategoryDropdownDTO, len(datas))

	for i, data := range datas {
		res[i] = dtoCat.CategoryDropdownDTO{
			ID:   data.ID,
			Name: s.cleanPath(data.Path),
		}
	}
	return res, nil
}

func (s *itemService) GetDropdownMerk(ctx context.Context, search string, page int) ([]dtoMerk.MerkDropdownDTO, error) {
	var res []dtoMerk.MerkDropdownDTO
	datas, err := s.mr.ReadDropdown(ctx, search, 10, page)
	if err != nil {
		return res, err
	}
	res = make([]dtoMerk.MerkDropdownDTO, len(datas))

	for i, data := range datas {
		res[i] = dtoMerk.MerkDropdownDTO{
			ID:   data.ID,
			Name: data.Name,
		}
	}
	return res, nil
}
func (s *itemService) cleanPath(path string) string {
	parts := strings.Split(path, "/")
	for j, p := range parts {
		if _, after, ok := strings.Cut(p, "-"); ok {
			parts[j] = after
		}
	}
	return strings.Join(parts, "/")
}
