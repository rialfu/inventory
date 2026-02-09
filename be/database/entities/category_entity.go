package entities

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100);not null" json:"name"`
	Depth     uint
	Path      string
	ParentID  *uint64    `gorm:"column:parent_id"`
	Parent    *Category  `gorm:"foreignKey:parent_id;references:ID"`
	Children  []Category `gorm:"foreignKey:parent_id;references:ID"`
	Items     []Item     `gorm:"foreignKey:category_id;references:ID"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	// IsVerified bool   `gorm:"default:false" json:"is_verified"`

	// Timestamp
}

func (m *Category) BeforeSave(tx *gorm.DB) (err error) {
	m.Name = strings.TrimSpace(m.Name)
	return
}
