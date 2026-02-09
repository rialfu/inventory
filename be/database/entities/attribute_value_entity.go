package entities

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type AttributeValue struct {
	ID              uint64                      `gorm:"primaryKey"`
	AttributeValue  string                      `gorm:"type:varchar(50);column:attribute_value;not null;" json:"attribute_value"`
	AttributeNameID uint64                      `gorm:"column:attribute_name_id"`
	AttributeName   AttributeName               `gorm:"foreignKey:attribute_name_id;references:ID"`
	CreatedAt       time.Time                   `gorm:"column:created_at"`
	UpdatedAt       time.Time                   `gorm:"column:updated_at"`
	Items           []ItemVariantAttributeValue `gorm:"foreignKey:attribute_value_id;references:ID"`
}

func (m *AttributeValue) BeforeSave(tx *gorm.DB) (err error) {
	m.AttributeValue = strings.TrimSpace(m.AttributeValue)
	return
}
