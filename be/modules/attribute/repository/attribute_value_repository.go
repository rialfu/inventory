package repository

import (
	"context"
	"errors"

	"rialfu/wallet/database/entities"
	"rialfu/wallet/pkg/helpers"

	"gorm.io/gorm"
)

type (
	AttributeValueRepository interface {
		Create(ctx context.Context, data entities.AttributeValue) (entities.AttributeValue, error)
		GetById(ctx context.Context, id string) (entities.AttributeValue, bool, error)
		Update(ctx context.Context, data entities.AttributeValue) (entities.AttributeValue, error)
		ReadAll(ctx context.Context, queryParam map[string][]string) ([]entities.AttributeValue, int, int, int64, error)
		ReadAllWithParent(ctx context.Context, queryParam map[string][]string, parentId string) ([]entities.AttributeValue, int, int, int64, error)
		GetByName(ctx context.Context, name string) (entities.AttributeValue, bool, error)
		GetByNameOrId(ctx context.Context, id string, name string) (entities.AttributeValue, bool, error)
		GetByIds(ctx context.Context, id []string) ([]entities.AttributeValue, error)
		GetByIdsWithAttrName(ctx context.Context, id []string) ([]entities.AttributeValue, error)
		GetByIdWithAttrName(ctx context.Context, id string) (entities.AttributeValue, bool, error)
	}

	attributeValueRepository struct {
		db *gorm.DB
	}
)

func NewAttributeValueRepository(db *gorm.DB) AttributeValueRepository {
	return &attributeValueRepository{
		db: db,
	}
}
func (r *attributeValueRepository) getDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("DB_TX").(*gorm.DB); ok {
		return tx
	}
	return r.db
}

func (r *attributeValueRepository) Create(ctx context.Context, data entities.AttributeValue) (entities.AttributeValue, error) {
	db := r.getDB(ctx)

	if err := db.WithContext(ctx).Create(&data).Error; err != nil {
		return entities.AttributeValue{}, err
	}
	var res entities.AttributeValue
	if err := db.Preload("AttributeName").First(&res, data.ID).Error; err != nil {
		return data, nil
	}
	return res, nil
}

func (r *attributeValueRepository) GetById(ctx context.Context, id string) (entities.AttributeValue, bool, error) {
	db := r.getDB(ctx)
	var data entities.AttributeValue
	if err := db.WithContext(ctx).Where("id = ?", id).Take(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.AttributeValue{}, false, nil
		}
		return entities.AttributeValue{}, false, err
	}

	return data, true, nil
}

func (r *attributeValueRepository) GetByIds(ctx context.Context, id []string) ([]entities.AttributeValue, error) {
	db := r.getDB(ctx)
	var datas []entities.AttributeValue
	if err := db.WithContext(ctx).Where("id in ?", id).Find(&datas).Error; err != nil {

		return datas, err
	}

	return datas, nil
}
func (r *attributeValueRepository) GetByIdsWithAttrName(ctx context.Context, id []string) ([]entities.AttributeValue, error) {
	db := r.getDB(ctx)
	var datas []entities.AttributeValue
	if err := db.WithContext(ctx).Preload("AttributeName").Where("id in ?", id).Find(&datas).Error; err != nil {

		return datas, err
	}

	return datas, nil
}
func (r *attributeValueRepository) GetByIdWithAttrName(ctx context.Context, id string) (entities.AttributeValue, bool, error) {
	db := r.getDB(ctx)
	var data entities.AttributeValue
	if err := db.WithContext(ctx).Where("id = ?", id).Preload("AttributeName").Take(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.AttributeValue{}, false, nil
		}
		return entities.AttributeValue{}, false, err
	}
	return data, true, nil
}

func (r *attributeValueRepository) GetByNameOrId(ctx context.Context, id string, name string) (entities.AttributeValue, bool, error) {
	db := r.getDB(ctx)
	var data entities.AttributeValue
	if err := db.WithContext(ctx).Where("id = ?", id).Or("attribute_value = ?", name).Take(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.AttributeValue{}, false, nil
		}
		return entities.AttributeValue{}, false, err
	}
	return data, true, nil
}
func (r *attributeValueRepository) GetByName(ctx context.Context, name string) (entities.AttributeValue, bool, error) {
	db := r.getDB(ctx)
	var data entities.AttributeValue
	if err := db.WithContext(ctx).Where("attribute_value = ?", name).Take(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.AttributeValue{}, false, nil
		}
		return entities.AttributeValue{}, false, err
	}
	return data, true, nil
}
func (r *attributeValueRepository) ReadAll(ctx context.Context, queryParam map[string][]string) ([]entities.AttributeValue, int, int, int64, error) {
	var datas []entities.AttributeValue
	db := r.db.Model(&entities.AttributeValue{})
	db, pagination := helpers.ApplyPagination(db, &entities.AttributeValue{}, queryParam)
	var total int64
	if err := db.Find(&datas).Error; err != nil {
		return []entities.AttributeValue{}, 0, 0, 0, err
	}
	if err := db.Limit(-1).Offset(-1).Count(&total).Error; err != nil {
		return []entities.AttributeValue{}, 0, 0, 0, err
	}
	return datas, pagination.Page, pagination.Limit, total, nil
}
func (r *attributeValueRepository) ReadAllWithParent(ctx context.Context, queryParam map[string][]string, parentId string) ([]entities.AttributeValue, int, int, int64, error) {
	var datas []entities.AttributeValue
	db := r.db.Model(&entities.AttributeValue{})
	// db = db.Preload("AttributeName")
	db = db.Where("attribute_name_id = ?", parentId)
	db, pagination := helpers.ApplyPagination(db, &entities.AttributeValue{}, queryParam)

	var total int64
	if err := db.Find(&datas).Error; err != nil {
		return []entities.AttributeValue{}, 0, 0, 0, err
	}
	if err := db.Limit(-1).Offset(-1).Count(&total).Error; err != nil {
		return []entities.AttributeValue{}, 0, 0, 0, err
	}
	return datas, pagination.Page, pagination.Limit, total, nil
}
func (r *attributeValueRepository) Update(ctx context.Context, data entities.AttributeValue) (entities.AttributeValue, error) {
	db := r.getDB(ctx)

	if err := db.WithContext(ctx).Updates(&data).Error; err != nil {
		return entities.AttributeValue{}, err
	}
	// var res entities.AttributeValue
	// if err := db.Preload("AttributeName").First(&res, data.ID).Error; err != nil {
	// 	return data, nil
	// }
	return data, nil
}
