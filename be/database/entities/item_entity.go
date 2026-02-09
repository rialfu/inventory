package entities

import "time"

type Item struct {
	ID         uint64 `gorm:"primaryKey"`
	Name       string
	CategoryID uint64 `gorm:"column:category_id"`
	Category   Category
	MerkID     uint64 `gorm:"column:merk_id"`
	Merk       Merk
	Variant    []ItemVariant `gorm:"foreignKey:item_id;references:ID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
