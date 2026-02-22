package dto

import "errors"

const (
// Failed
// MESSAGE_FAILED_REGISTER_USER = "failed create user"

// Success
// MESSAGE_SUCCESS_REGISTER_USER           = "success create user"
)

var (
	ErrLockedEdit       = errors.New("locked to edit")
	ErrMerkNotFound     = errors.New("Merk Not Found")
	ErrCategoryNotFound = errors.New("Category Not Found")
)

type (
	ItemCreateRequest struct {
		ID       *string `json:"id,omitempty"`
		Name     string  `json:"name" form:"name" validate:"required,min=2,max=50"`
		Category uint64  `json:"category" validate:"required"`
		Merk     uint64  `json:"merk" validate:"required"`
		// Variant  []ItemVariantCreateRequest `json:"variants" binding:"required,dive"`
	}
	ItemVariantCreateRequest struct {
		ID          *uint64  `json:"id,omitempty"`
		SKU         string   `json:"sku" validate:"sku"`
		Description string   `json:"description" validate:"description"`
		Item        *string  `json:"item,omitempty"`
		Attributes  []uint64 `json:"attributes"`
	}

	ItemResponse struct {
		ID           uint64 `json:"id"`
		Name         string `json:"name"`
		CategoryName string `json:"category_name"`
		CategoryID   uint64 `json:"category_id"`
		MerkName     string `json:"merk_name"`
		MerkID       uint64 `json:"merk_id"`
	}
	VariantResponse struct {
		ID          uint64                  `json:"id"`
		SKU         string                  `json:"sku"`
		Description string                  `json:"description"`
		Attributes  []AttributeNameResponse `json:"attributes"`
	}
	AttributeNameResponse struct {
		ID    *uint64                  `json:"id,omitempty"`
		Name  string                   `json:"name"`
		Value []AttributeValueResponse `json:"value"`
	}
	AttributeValueResponse struct {
		ID   *uint64 `json:"id,omitempty"`
		Name string  `json:"name"`
	}
)
