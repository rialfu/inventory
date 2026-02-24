package repository

import (
	"context"
	"errors"

	"rialfu/wallet/database/entities"
	"rialfu/wallet/pkg/helpers"

	"gorm.io/gorm"
)

type (
	AttributeNameRepository interface {
		Create(ctx context.Context, data entities.AttributeName) (entities.AttributeName, error)
		GetById(ctx context.Context, id string) (entities.AttributeName, bool, error)
		Update(ctx context.Context, data entities.AttributeName) (entities.AttributeName, error)
		ReadAll(ctx context.Context, queryParam map[string][]string) ([]entities.AttributeName, int, int, int64, error)
		GetByName(ctx context.Context, name string) (entities.AttributeName, bool, error)
		GetByNameOrId(ctx context.Context, id string, name string) (entities.AttributeName, bool, error)
		ReadDropdown(ctx context.Context, search string, limit int, page int) ([]entities.AttributeName, error)
	}

	attributeNameRepository struct {
		db *gorm.DB
	}
)

func NewAttributeNameRepository(db *gorm.DB) AttributeNameRepository {
	return &attributeNameRepository{
		db: db,
	}
}
func (r *attributeNameRepository) getDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("DB_TX").(*gorm.DB); ok {
		return tx
	}
	return r.db
}

func (r *attributeNameRepository) Create(ctx context.Context, data entities.AttributeName) (entities.AttributeName, error) {
	db := r.getDB(ctx)

	if err := db.WithContext(ctx).Create(&data).Error; err != nil {
		return entities.AttributeName{}, err
	}

	return data, nil
}

func (r *attributeNameRepository) GetById(ctx context.Context, id string) (entities.AttributeName, bool, error) {
	db := r.getDB(ctx)
	var data entities.AttributeName
	if err := db.WithContext(ctx).Where("id = ?", id).Take(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.AttributeName{}, false, nil
		}
		return entities.AttributeName{}, false, err
	}

	return data, true, nil
}

func (r *attributeNameRepository) GetByNameOrId(ctx context.Context, id string, name string) (entities.AttributeName, bool, error) {
	db := r.getDB(ctx)
	var data entities.AttributeName
	if err := db.WithContext(ctx).Where("id = ?", id).Or("attribute_name = ?", name).Take(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.AttributeName{}, false, nil
		}
		return entities.AttributeName{}, false, err
	}
	return data, true, nil
}
func (r *attributeNameRepository) GetByName(ctx context.Context, name string) (entities.AttributeName, bool, error) {
	db := r.getDB(ctx)
	var data entities.AttributeName
	if err := db.WithContext(ctx).Where("attribute_name = ?", name).Take(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.AttributeName{}, false, nil
		}
		return entities.AttributeName{}, false, err
	}
	return data, true, nil
}
func (r *attributeNameRepository) ReadAll(ctx context.Context, queryParam map[string][]string) ([]entities.AttributeName, int, int, int64, error) {
	var datas []entities.AttributeName
	db := r.db.Model(&entities.AttributeName{})
	db, pagination := helpers.ApplyPagination(db, &entities.AttributeName{}, queryParam)
	var total int64
	if err := db.Find(&datas).Error; err != nil {
		return []entities.AttributeName{}, 0, 0, 0, err
	}
	if err := db.Limit(-1).Offset(-1).Count(&total).Error; err != nil {
		return []entities.AttributeName{}, 0, 0, 0, err
	}
	return datas, pagination.Page, pagination.Limit, total, nil
}

func (r *attributeNameRepository) Update(ctx context.Context, data entities.AttributeName) (entities.AttributeName, error) {
	db := r.getDB(ctx)

	if err := db.WithContext(ctx).Updates(&data).Error; err != nil {
		return entities.AttributeName{}, err
	}

	return data, nil
}
func (r *attributeNameRepository) ReadDropdown(ctx context.Context, search string, limit int, page int) ([]entities.AttributeName, error) {
	db := r.getDB(ctx)
	if search != "" {
		db = db.Where("attribute_name ILIKE ?", "%"+search+"%")
	}
	if limit > 100 {
		limit = 100
	} else if limit < 1 {
		limit = 10
	}
	if page < 1 {
		page = 1
	}

	db = db.Order("attribute_name asc")
	var data []entities.AttributeName
	offset := (page - 1) * limit
	if err := db.WithContext(ctx).Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}
