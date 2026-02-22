package entities

import "time"

type ItemVariant struct {
	ID       uint64 `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(100);"`
	ItemID   uint64 `gorm:"column:item_id"`
	SKU      string `gorm:"unique;not null"`
	ImageURL string `gorm:"size:500"`
	Stock    int32  `gorm:"default:0"`

	Description string `gorm:"type:text;"`

	Item Item                        `gorm:"foreignKey:item_id;references:ID"`
	Attr []ItemVariantAttributeValue `gorm:"foreignKey:item_variant_id;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
