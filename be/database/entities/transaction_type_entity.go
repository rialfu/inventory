package entities

type TransactionType struct {
	ID uint64 `gorm:"primaryKey"`

	TypeName string `gorm:"size:20;unique;not null;"`
}
