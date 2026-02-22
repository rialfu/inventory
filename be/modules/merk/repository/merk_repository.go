package repository

import (
	"context"
	"errors"

	"rialfu/wallet/database/entities"
	"rialfu/wallet/pkg/helpers"

	"gorm.io/gorm"
)

type (
	MerkRepository interface {
		Create(ctx context.Context, data entities.Merk) (entities.Merk, error)
		GetById(ctx context.Context, id string) (entities.Merk, bool, error)
		Update(ctx context.Context, data entities.Merk) (entities.Merk, error)
		ReadAll(ctx context.Context, queryParam map[string][]string) ([]entities.Merk, int, int, int64, error)
		GetByName(ctx context.Context, name string) (entities.Merk, bool, error)
		GetByNameOrId(ctx context.Context, id string, name string) (entities.Merk, bool, error)
		ReadDropdown(ctx context.Context, search string, limit int, page int) ([]entities.Merk, error)
	}

	merkRepository struct {
		db *gorm.DB
	}
)

func NewMerkRepository(db *gorm.DB) MerkRepository {
	return &merkRepository{
		db: db,
	}
}
func (r *merkRepository) getDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("DB_TX").(*gorm.DB); ok {
		return tx
	}
	return r.db
}

func (r *merkRepository) Create(ctx context.Context, data entities.Merk) (entities.Merk, error) {
	db := r.getDB(ctx)

	if err := db.WithContext(ctx).Create(&data).Error; err != nil {
		return entities.Merk{}, err
	}

	return data, nil
}

func (r *merkRepository) GetById(ctx context.Context, id string) (entities.Merk, bool, error) {
	db := r.getDB(ctx)
	var data entities.Merk
	if err := db.WithContext(ctx).Where("id = ?", id).Take(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Merk{}, false, nil
		}
		return entities.Merk{}, false, err
	}

	return data, true, nil
}

func (r *merkRepository) GetByNameOrId(ctx context.Context, id string, name string) (entities.Merk, bool, error) {
	db := r.getDB(ctx)
	var data entities.Merk
	if err := db.WithContext(ctx).Where("id = ?", id).Or("name = ?", name).Take(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Merk{}, false, nil
		}
		return entities.Merk{}, false, err
	}
	return data, true, nil
}
func (r *merkRepository) GetByName(ctx context.Context, name string) (entities.Merk, bool, error) {
	db := r.getDB(ctx)
	var data entities.Merk
	if err := db.WithContext(ctx).Where("name = ?", name).Take(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Merk{}, false, nil
		}
		return entities.Merk{}, false, err
	}
	return data, true, nil
}
func (r *merkRepository) ReadAll(ctx context.Context, queryParam map[string][]string) ([]entities.Merk, int, int, int64, error) {
	var datas []entities.Merk
	db := r.db.Model(&entities.Merk{})
	db, pagination := helpers.ApplyPagination(db, &entities.Merk{}, queryParam)
	var total int64
	if err := db.Find(&datas).Error; err != nil {
		return []entities.Merk{}, 0, 0, 0, err
	}
	if err := db.Limit(-1).Offset(-1).Count(&total).Error; err != nil {
		return []entities.Merk{}, 0, 0, 0, err
	}
	return datas, pagination.Page, pagination.Limit, total, nil
}

func (r *merkRepository) Update(ctx context.Context, data entities.Merk) (entities.Merk, error) {
	db := r.getDB(ctx)

	if err := db.WithContext(ctx).Updates(&data).Error; err != nil {
		return entities.Merk{}, err
	}

	return data, nil
}
func (r *merkRepository) ReadDropdown(ctx context.Context, search string, limit int, page int) ([]entities.Merk, error) {
	db := r.getDB(ctx)
	if search != "" {
		db = db.Where("name ILIKE ?", "%"+search+"%")
	}
	if limit > 100 {
		limit = 100
	} else if limit < 1 {
		limit = 10
	}
	if page < 1 {
		page = 1
	}

	db = db.Order("name asc")
	var data []entities.Merk
	offset := (page - 1) * limit
	if err := db.WithContext(ctx).Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}
