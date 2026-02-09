package repository

import (
	"context"
	"errors"
	"fmt"
	"rialfu/wallet/database/entities"
	"rialfu/wallet/modules/category/dto"
	"rialfu/wallet/pkg/helpers"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	CategoryRepository interface {
		Create(ctx context.Context, data entities.Category, parentID uint64) (entities.Category, error)
		GetById(ctx context.Context, id string, isLock bool) (entities.Category, bool, error)
		ReadDropdown(ctx context.Context, search string, limit int, page int) ([]entities.Category, error)
		UpdateSubtreeDepth(ctx context.Context, oldDepth, newDepth int, path string) error
		UpdateSubtreePath(ctx context.Context, oldPath, newPath string) error
		Update(ctx context.Context, id string, data map[string]any) error
		ReadAll(ctx context.Context, queryParam map[string][]string, parentId *string) ([]entities.Category, int, int, int64, error)
	}

	categoryRepository struct {
		db *gorm.DB
	}
)

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}
func (r *categoryRepository) getDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("DB_TX").(*gorm.DB); ok {
		return tx
	}
	return r.db
}

func (r *categoryRepository) Create(ctx context.Context, data entities.Category, parentID uint64) (entities.Category, error) {
	db := r.getDB(ctx)
	tx := db.Begin()
	if tx.Error != nil {
		return entities.Category{}, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()
	var parent entities.Category
	if parentID != 0 {
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).WithContext(ctx).First(&parent, parentID).Error; err != nil {
			tx.Rollback()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return entities.Category{}, dto.ErrGetParentNotFound
			}
			return entities.Category{}, err
		}
		data.ParentID = &parent.ID
		data.Depth = parent.Depth + 1

	}

	if err := tx.WithContext(ctx).Create(&data).Error; err != nil {
		tx.Rollback()
		return entities.Category{}, err
	}
	if parentID != 0 {
		data.Path = fmt.Sprintf("%s/%d-%s", parent.Path, data.ID, data.Name)

	} else {
		data.Path = fmt.Sprintf("%d-%s", data.ID, data.Name)
	}
	if err := tx.WithContext(ctx).Save(&data).Error; err != nil {
		tx.Rollback()
		return entities.Category{}, err
	}
	if err := tx.Commit().Error; err != nil {
		return entities.Category{}, err
	}
	return data, nil
}

func (r *categoryRepository) GetById(ctx context.Context, id string, isLock bool) (entities.Category, bool, error) {
	var data entities.Category
	db := r.getDB(ctx)
	if isLock {
		db = db.Clauses(clause.Locking{Strength: "UPDATE"})
	}
	if err := db.WithContext(ctx).Where("id = ?", id).Take(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Category{}, false, nil
		}
		return entities.Category{}, false, err
	}
	return data, true, nil
}

func (r *categoryRepository) Update(ctx context.Context, id string, data map[string]any) error {
	db := r.getDB(ctx)
	return db.WithContext(ctx).Model(&entities.Category{}).Where("id = ?", id).Updates(data).Error

}

func (r *categoryRepository) UpdateSubtreeDepth(ctx context.Context, oldDepth, newDepth int, path string) error {
	db := r.getDB(ctx)

	diff := newDepth - oldDepth
	return db.WithContext(ctx).
		Exec(`
            UPDATE categories
            SET depth = depth + ?
            WHERE path LIKE ?
        `,
			diff,
			path+"%",
		).Error
}
func (r *categoryRepository) UpdateSubtreePath(ctx context.Context, oldPath, newPath string) error {
	db := r.getDB(ctx)
	return db.WithContext(ctx).
		Exec(`
            UPDATE categories
            SET path = CONCAT(?, SUBSTRING(path, ?))
            WHERE path LIKE ?
        `,
			newPath,
			len(oldPath)+1,
			oldPath+"%",
		).Error
}

func (r *categoryRepository) ReadAll(ctx context.Context, queryParam map[string][]string, parentId *string) ([]entities.Category, int, int, int64, error) {
	var data []entities.Category
	db := r.db.Model(&entities.Category{})
	db, pagination := helpers.ApplyPagination(db, &entities.Category{}, queryParam)
	var total int64
	if parentId != nil {
		db = db.Where("parent_id = ?", parentId)
	} else {
		// fmt.Println("jalan kosong")
		db = db.Where("parent_id IS NULL")
	}
	if err := db.Find(&data).Error; err != nil {
		return []entities.Category{}, 0, 0, 0, err
	}
	if err := db.Limit(-1).Offset(-1).Count(&total).Error; err != nil {
		return []entities.Category{}, 0, 0, 0, err
	}
	return data, pagination.Page, pagination.Limit, total, nil

}
func (r *categoryRepository) ReadDropdown(ctx context.Context, search string, limit int, page int) ([]entities.Category, error) {
	db := r.getDB(ctx)
	if search != "" {
		db = db.Where("name LIKE ?", "%"+search+"%")
	}
	if limit > 100 {
		limit = 100
	} else if limit < 1 {
		limit = 10
	}
	if page < 1 {
		page = 1
	}

	db = db.Order("depth asc").Order("parent_id NULLS FIRST").Order("name asc")
	var data []entities.Category
	offset := (page - 1) * limit
	if err := db.WithContext(ctx).Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil

	// var result dto.CategoryDropdownDTO
	//
	// if err := db.Raw(`
	// 	WITH RECURSIVE tree AS (
	// 		SELECT
	// 			id,
	// 			name,
	// 			parent_id,
	// 			name::text AS path,
	// 			0 AS depth
	// 		FROM categories
	// 		WHERE parent_id IS NULL

	// 		UNION ALL

	// 		SELECT
	// 			c.id,
	// 			c.name,
	// 			c.parent_id,
	// 			t.path || '/' || c.name,
	// 			t.depth + 1
	// 		FROM categories c
	// 		JOIN tree t
	// 			ON c.parent_id = t.id
	// 	)
	// 	SELECT
	// 		id,
	// 		name,
	// 		path,
	// 		depth
	// 	FROM tree
	// 	WHERE path ILIKE '%' ? '%'
	// 	ORDER BY path ASC
	// 	LIMIT ? OFFSET ?;
	// 	`, search, limit, offset).Scan(&result).Error; err != nil {
	// 	return result, err
	// }
	// return result, nil
}

// WITH RECURSIVE tree AS (
//     SELECT
//         id,
//         name,
//         parent_id,
//         name::text AS path,
//         0 AS depth
//     FROM categories
//     WHERE parent_id IS NULL

//     UNION ALL

//     SELECT
//         c.id,
//         c.name,
//         c.parent_id,
//         t.path || '/' || c.name,
//         t.depth + 1
//     FROM categories c
//     JOIN tree t
//         ON c.parent_id = t.id
// )
// SELECT
//     id,
//     name,
//     path,
//     depth
// FROM tree
// WHERE name ILIKE '%' || :keyword || '%'
// ORDER BY name ASC
// LIMIT :limit OFFSET :offset;
