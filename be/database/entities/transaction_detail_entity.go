package entities

import "github.com/shopspring/decimal"

type TransactionDetail struct {
	ID            uint64
	TransactionID uint64 `gorm:"column:transaction_history_id;not null;"`
	VariantItemID uint64 `gorm:"column:variant_item_id;not null;"`
	Qty           uint32
	Price         decimal.Decimal `gorm:"type:numeric(16,2);not null"`
	/*
		IN  → harga modal saat masuk

		OUT → harga jual saat keluar
	*/
	Transaction TransactionHistory `gorm:"foreignKey:transaction_history_id;references:ID"`
	ItemVariant ItemVariant        `gorm:"foreignKey:variant_item_id;references:ID"`
}
