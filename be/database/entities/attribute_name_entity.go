package entities

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type AttributeName struct {
	ID              uint64           `gorm:"primaryKey"`
	AttributeName   string           `gorm:"type:varchar(50);column:attribute_name;not null;unique;" json:"attribute_name"`
	AttributeValues []AttributeValue `gorm:"foreignKey:attribute_name_id;references:ID"`
	CreatedAt       time.Time        `gorm:"column:created_at"`
	UpdatedAt       time.Time        `gorm:"column:updated_at"`
}

func (m *AttributeName) BeforeSave(tx *gorm.DB) (err error) {
	m.AttributeName = strings.TrimSpace(m.AttributeName)
	return
}
