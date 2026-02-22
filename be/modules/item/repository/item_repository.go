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
	ItemRepository interface {
		Create(ctx context.Context, data entities.Item) (entities.Item, error)
		GetById(ctx context.Context, id string, isLock bool) (entities.Item, bool, error)
		Update(ctx context.Context, data entities.Item) (entities.Item, error)
		ReadAll(ctx context.Context, queryParam map[string][]string) ([]entities.Item, int, int, int64, error)
		GetByName(ctx context.Context, name string) (entities.Item, bool, error)
		GetByNameOrId(ctx context.Context, id string, name string) (entities.Item, bool, error)
		GetByIdWithMerkAndCategory(ctx context.Context, id string) (entities.Item, bool, error)
	}

	itemRepository struct {
		db *gorm.DB
	}
)

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{
		db: db,
	}
}
func (r *itemRepository) getDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("DB_TX").(*gorm.DB); ok {
		return tx
	}
	return r.db
}

func (r *itemRepository) Create(ctx context.Context, data entities.Item) (entities.Item, error) {
	db := r.getDB(ctx)

	if err := db.WithContext(ctx).Create(&data).Error; err != nil {
		return entities.Item{}, err
	}

	return data, nil
}
func (r *itemRepository) Update(ctx context.Context, data entities.Item) (entities.Item, error) {
	db := r.getDB(ctx)

	if err := db.WithContext(ctx).Updates(&data).Error; err != nil {
		return entities.Item{}, err
	}

	return data, nil
}

func (r *itemRepository) GetById(ctx context.Context, id string, isLock bool) (entities.Item, bool, error) {
	db := r.getDB(ctx)
	var data entities.Item
	if isLock {
		db = db.Clauses(clause.Locking{Strength: "UPDATE"})
	}
	if err := db.WithContext(ctx).Where("id = ?", id).Take(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Item{}, false, nil
		}
		return entities.Item{}, false, err
	}

	return data, true, nil
}
func (r *itemRepository) GetByIdWithMerkAndCategory(ctx context.Context, id string) (entities.Item, bool, error) {
	db := r.getDB(ctx)
	var data entities.Item

	db = db.Preload("Merk").Preload("Category")
	if err := db.WithContext(ctx).Where("id = ?", id).Take(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Item{}, false, nil
		}
		return entities.Item{}, false, err
	}

	return data, true, nil
}

func (r *itemRepository) GetByNameOrId(ctx context.Context, id string, name string) (entities.Item, bool, error) {
	db := r.getDB(ctx)
	var data entities.Item
	if err := db.WithContext(ctx).Where("id = ?", id).Or("name = ?", name).Take(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Item{}, false, nil
		}
		return entities.Item{}, false, err
	}
	return data, true, nil
}
func (r *itemRepository) GetByName(ctx context.Context, name string) (entities.Item, bool, error) {
	db := r.getDB(ctx)
	var data entities.Item
	if err := db.WithContext(ctx).Where("name = ?", name).Take(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Item{}, false, nil
		}
		return entities.Item{}, false, err
	}
	return data, true, nil
}
func (r *itemRepository) ReadAll(ctx context.Context, queryParam map[string][]string) ([]entities.Item, int, int, int64, error) {
	var datas []entities.Item
	db := r.db.Model(&entities.Item{})
	db = db.Preload("Merk").Preload("Category")

	if vals, ok := queryParam["merk"]; ok && len(vals) > 0 {
		db = db.Where("merk_id = ?", vals[0])
	}
	if vals, ok := queryParam["category"]; ok && len(vals) > 0 {
		db = db.Where("category_id = ?", vals[0])
	}

	db, pagination := helpers.ApplyPagination(db, &entities.Item{}, queryParam)
	var total int64
	if err := db.Find(&datas).Error; err != nil {
		return []entities.Item{}, 0, 0, 0, err
	}
	if err := db.Limit(-1).Offset(-1).Count(&total).Error; err != nil {
		return []entities.Item{}, 0, 0, 0, err
	}
	return datas, pagination.Page, pagination.Limit, total, nil
}
