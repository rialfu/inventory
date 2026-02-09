package entities

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type Merk struct {
	ID        uint64    `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(50);not null;unique;" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Items     []Item    `gorm:"foreignKey:merk_id;references:ID"`

	// Timestamp
}

func (m *Merk) BeforeSave(tx *gorm.DB) (err error) {
	m.Name = strings.TrimSpace(m.Name)
	return
}
