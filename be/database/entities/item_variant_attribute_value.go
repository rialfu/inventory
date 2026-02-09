package entities

import (
	"time"
)

type ItemVariantAttributeValue struct {
	ID               uint64         `gorm:"primaryKey"`
	ItemVariantID    uint64         `gorm:"column:item_variant_id"`
	ItemVariant      ItemVariant    `gorm:"foreignKey:item_variant_id;references:ID"`
	AttributeValueID uint64         `gorm:"column:attribute_value_id"`
	AttributeValue   AttributeValue `gorm:"foreignKey:attribute_value_id;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
