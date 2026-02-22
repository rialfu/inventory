package entities

import "time"

type TransactionHistory struct {
	ID        uint64    `gorm:"primaryKey"`
	Code      string    `gorm:"size:50;uniqueIndex"`
	TypeID    uint64    `gorm:"column:transaction_type_id;not null;"`
	DateTrans time.Time `gorm:"type:date;not null;"`
	Note      string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"column:created_at"`

	TransactionType TransactionType `gorm:"foreignKey:transaction_type_id;references:ID"`
}
