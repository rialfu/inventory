package service

import (
	"context"
	"rialfu/wallet/database/entities"
	"rialfu/wallet/modules/attribute/dto"
	"rialfu/wallet/modules/attribute/repository"
	"rialfu/wallet/pkg/constants"
	"rialfu/wallet/pkg/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AttributeService interface {
	CreateAttributeName(ctx context.Context, req dto.AttributeNameCreateRequest) (dto.AttributeNameResponse, error)
	CreateAttributeValue(ctx context.Context, req dto.AttributeValueCreateRequest) (dto.AttributeValueResponse, error)
	UpdateAttributeName(ctx context.Context, req dto.AttributeNameUpdateRequest, id string) (dto.AttributeNameResponse, error)
	UpdateAttributeValue(ctx context.Context, req dto.AttributeValueUpdateRequest, id string) (dto.AttributeValueResponse, error)
	GetAllAttributeName(ctx *gin.Context) (helpers.PaginateData[dto.AttributeNameResponse], error)
	GetAllAttributeValueBasedParent(ctx *gin.Context, parentId string) (dto.AttributeNameWithValueResponse, error)
	GetAttributeName(ctx context.Context, id string) (dto.AttributeNameResponse, error)
}

type attributeService struct {
	anrepository repository.AttributeNameRepository
	avrepository repository.AttributeValueRepository
	db           *gorm.DB
}

func NewAttributeService(
	anr repository.AttributeNameRepository,
	avr repository.AttributeValueRepository,
	db *gorm.DB,
) AttributeService {
	return &attributeService{
		anrepository: anr,
		avrepository: avr,
		db:           db,
	}
}

func (s *attributeService) CreateAttributeName(ctx context.Context, req dto.AttributeNameCreateRequest) (dto.AttributeNameResponse, error) {
	_, isExist, err := s.anrepository.GetByName(ctx, req.Name)

	if err != nil {
		return dto.AttributeNameResponse{}, err
	}
	if isExist {
		return dto.AttributeNameResponse{}, constants.ErrValueNotUniq
	}
	data := entities.AttributeName{
		AttributeName: req.Name,
	}
	updated, err := s.anrepository.Create(ctx, data)
	res := dto.AttributeNameResponse{
		ID:   updated.ID,
		Name: updated.AttributeName,
	}

	return res, err
}
func (s *attributeService) GetAttributeName(ctx context.Context, id string) (dto.AttributeNameResponse, error) {
	data, isExist, err := s.anrepository.GetById(ctx, id)

	if err != nil {
		return dto.AttributeNameResponse{}, err
	}
	if isExist == false {
		return dto.AttributeNameResponse{}, constants.ErrDataNotFound
	}

	res := dto.AttributeNameResponse{
		ID:   data.ID,
		Name: data.AttributeName,
	}

	return res, err
}

func (s *attributeService) CreateAttributeValue(ctx context.Context, req dto.AttributeValueCreateRequest) (dto.AttributeValueResponse, error) {
	an, isExist, err := s.anrepository.GetById(ctx, req.AttributeNameID)
	if err != nil {
		return dto.AttributeValueResponse{}, err
	}
	if isExist == false {
		return dto.AttributeValueResponse{}, constants.ErrDataNotFound
	}
	data := entities.AttributeValue{
		AttributeValue:  req.Name,
		AttributeNameID: an.ID,
	}
	updated, err := s.avrepository.Create(ctx, data)
	res := dto.AttributeValueResponse{
		ID:              updated.ID,
		Name:            updated.AttributeValue,
		AttributeName:   &updated.AttributeName.AttributeName,
		AttributeNameID: &updated.AttributeName.ID,
	}

	return res, err
}

func (s *attributeService) GetAllAttributeName(ctx *gin.Context) (helpers.PaginateData[dto.AttributeNameResponse], error) {
	datas, page, limit, total, err := s.anrepository.ReadAll(ctx, ctx.Request.URL.Query())

	if err != nil {
		return helpers.PaginateData[dto.AttributeNameResponse]{}, err
	}
	results := make([]dto.AttributeNameResponse, 0, len(datas))
	for _, u := range datas {
		item := dto.AttributeNameResponse{
			ID:   u.ID,
			Name: u.AttributeName,
		}
		results = append(results, item)
	}

	return helpers.PaginateData[dto.AttributeNameResponse]{Data: results, Limit: limit, Page: page, Total: total}, nil
}
func (s *attributeService) GetAllAttributeValueBasedParent(ctx *gin.Context, parentId string) (dto.AttributeNameWithValueResponse, error) {
	var res dto.AttributeNameWithValueResponse
	mainData, isExist, err := s.anrepository.GetById(ctx, parentId)
	if err != nil {
		return res, err
	}
	if isExist == false {
		return res, constants.ErrDataNotFound
	}
	datas, page, limit, total, err := s.avrepository.ReadAllWithParent(ctx, ctx.Request.URL.Query(), parentId)

	if err != nil {
		return res, err
	}
	res.ID = mainData.ID
	res.Name = mainData.AttributeName

	results := make([]dto.AttributeValueResponse, 0, len(datas))
	for _, u := range datas {
		item := dto.AttributeValueResponse{
			ID:   u.ID,
			Name: u.AttributeValue,
		}
		results = append(results, item)
	}
	res.Values = helpers.PaginateData[dto.AttributeValueResponse]{Data: results, Limit: limit, Page: page, Total: total}
	return res, nil
}

func (s *attributeService) UpdateAttributeName(ctx context.Context, req dto.AttributeNameUpdateRequest, id string) (dto.AttributeNameResponse, error) {
	data, isExist, err := s.anrepository.GetByNameOrId(ctx, id, req.Name)
	if err != nil {
		return dto.AttributeNameResponse{}, err
	}
	if isExist == false {
		if strconv.FormatUint(data.ID, 10) != id {
			return dto.AttributeNameResponse{}, constants.ErrValueNotUniq
		}

	}

	data.AttributeName = req.Name

	updated, err := s.anrepository.Update(ctx, data)
	res := dto.AttributeNameResponse{
		ID:   updated.ID,
		Name: updated.AttributeName,
	}

	return res, err
}
func (s *attributeService) UpdateAttributeValue(ctx context.Context, req dto.AttributeValueUpdateRequest, id string) (dto.AttributeValueResponse, error) {
	data, isExist, err := s.avrepository.GetById(ctx, id)
	if err != nil {
		return dto.AttributeValueResponse{}, err
	}
	if isExist == false {
		if strconv.FormatUint(data.ID, 10) != id {
			return dto.AttributeValueResponse{}, constants.ErrValueNotUniq
		}

	}
	an, isExist, err := s.anrepository.GetById(ctx, req.AttributeNameID)
	if err != nil {
		return dto.AttributeValueResponse{}, err
	}
	if isExist == false {
		return dto.AttributeValueResponse{}, constants.ErrDataNotFound
	}
	data.AttributeValue = req.Name
	data.AttributeNameID = an.ID

	updated, err := s.avrepository.Update(ctx, data)
	res := dto.AttributeValueResponse{
		ID:              updated.ID,
		Name:            updated.AttributeValue,
		AttributeName:   &an.AttributeName,
		AttributeNameID: &an.ID,
	}

	return res, err
}
