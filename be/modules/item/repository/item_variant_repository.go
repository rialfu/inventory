package repository

import (
	"context"
	"errors"

	"rialfu/wallet/database/entities"
	"rialfu/wallet/pkg/helpers"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	ItemVariantRepository interface {
		Create(ctx context.Context, data entities.ItemVariant) (entities.ItemVariant, error)
		GetById(ctx context.Context, id string, isLock bool) (entities.ItemVariant, bool, error)
		Update(ctx context.Context, data entities.ItemVariant) (entities.ItemVariant, error)
		ReadAll(ctx context.Context, queryParam map[string][]string) ([]entities.ItemVariant, int, int, int64, error)
		GetByName(ctx context.Context, name string) (entities.ItemVariant, bool, error)
		GetByNameOrId(ctx context.Context, id string, name string) (entities.ItemVariant, bool, error)
		GetBySKUs(ctx context.Context, sku []string, isLock bool) ([]entities.ItemVariant, error)
		GetByIdWithAttribute(ctx context.Context, id string) (entities.ItemVariant, bool, error)
		SyncVariantAttribute(ctx context.Context, datas []entities.ItemVariantAttributeValue, variantID uint64) error
	}

	itemVariantRepository struct {
		db *gorm.DB
	}
)

func NewItemVariantRepository(db *gorm.DB) ItemVariantRepository {
	return &itemVariantRepository{
		db: db,
	}
}
func (r *itemVariantRepository) getDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("DB_TX").(*gorm.DB); ok {
		return tx
	}
	return r.db
}

func (r *itemVariantRepository) Create(ctx context.Context, data entities.ItemVariant) (entities.ItemVariant, error) {
	db := r.getDB(ctx)

	if err := db.WithContext(ctx).Create(&data).Error; err != nil {
		return entities.ItemVariant{}, err
	}

	return data, nil
}

func (r *itemVariantRepository) GetById(ctx context.Context, id string, isLock bool) (entities.ItemVariant, bool, error) {
	db := r.getDB(ctx)
	var data entities.ItemVariant
	if isLock {
		db = db.Clauses(clause.Locking{Strength: "SHARE"})
	}
	if err := db.WithContext(ctx).Where("id = ?", id).Take(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.ItemVariant{}, false, nil
		}
		return entities.ItemVariant{}, false, err
	}

	return data, true, nil
}

func (r *itemVariantRepository) GetByIdWithAttribute(ctx context.Context, id string) (entities.ItemVariant, bool, error) {
	db := r.getDB(ctx)
	var data entities.ItemVariant
	db = db.Preload("Attr").Preload("Attr.AttributeValue").Preload("Attr.AttributeValue.AttributeName")
	if err := db.WithContext(ctx).Where("id = ?", id).Take(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.ItemVariant{}, false, nil
		}
		return entities.ItemVariant{}, false, err
	}

	return data, true, nil
}

func (r *itemVariantRepository) GetByNameOrId(ctx context.Context, id string, name string) (entities.ItemVariant, bool, error) {
	db := r.getDB(ctx)
	var data entities.ItemVariant
	if err := db.WithContext(ctx).Where("id = ?", id).Or("name = ?", name).Take(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.ItemVariant{}, false, nil
		}
		return entities.ItemVariant{}, false, err
	}
	return data, true, nil
}
func (r *itemVariantRepository) GetByName(ctx context.Context, name string) (entities.ItemVariant, bool, error) {
	db := r.getDB(ctx)
	var data entities.ItemVariant
	if err := db.WithContext(ctx).Where("name = ?", name).Take(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.ItemVariant{}, false, nil
		}
		return entities.ItemVariant{}, false, err
	}
	return data, true, nil
}
func (r *itemVariantRepository) GetBySKUs(ctx context.Context, sku []string, isLock bool) ([]entities.ItemVariant, error) {
	db := r.getDB(ctx)
	var datas []entities.ItemVariant
	if isLock {
		db = db.Clauses(clause.Locking{Strength: "SHARE"})
	}
	if err := db.WithContext(ctx).Where("sku = ?", sku).Find(&datas).Error; err != nil {
		return datas, err
	}
	return datas, nil
}
func (r *itemVariantRepository) ReadAll(ctx context.Context, queryParam map[string][]string) ([]entities.ItemVariant, int, int, int64, error) {
	var datas []entities.ItemVariant
	db := r.db.Model(&entities.ItemVariant{})
	db, pagination := helpers.ApplyPagination(db, &entities.ItemVariant{}, queryParam)
	var total int64
	if err := db.Find(&datas).Error; err != nil {
		return []entities.ItemVariant{}, 0, 0, 0, err
	}
	if err := db.Limit(-1).Offset(-1).Count(&total).Error; err != nil {
		return []entities.ItemVariant{}, 0, 0, 0, err
	}
	return datas, pagination.Page, pagination.Limit, total, nil
}

func (r *itemVariantRepository) Update(ctx context.Context, data entities.ItemVariant) (entities.ItemVariant, error) {
	db := r.getDB(ctx)

	if err := db.WithContext(ctx).Updates(&data).Error; err != nil {
		return entities.ItemVariant{}, err
	}

	return data, nil
}
func (r itemVariantRepository) SyncVariantAttribute(ctx context.Context, datas []entities.ItemVariantAttributeValue, variantID uint64) error {
	db := r.getDB(ctx)

	if err := db.WithContext(ctx).Where("item_variant_id = ?", variantID).
		Delete(&entities.ItemVariantAttributeValue{}).Error; err != nil {
		return err
	}
	db = r.getDB(ctx)
	if err := db.WithContext(ctx).Create(&datas).Error; err != nil {
		return err
	}
	return nil
}
