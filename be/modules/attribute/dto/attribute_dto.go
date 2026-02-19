package dto

import "rialfu/wallet/pkg/helpers"

const (
// Failed
// MESSAGE_FAILED_REGISTER_USER = "failed create user"

// Success
// MESSAGE_SUCCESS_REGISTER_USER           = "success create user"
)

var (
// ErrGetUserById        = errors.New("failed to get user by id")
)

type (
	AttributeNameCreateRequest struct {
		Name string `json:"name" form:"name" validate:"required,min=2,max=50"`
	}

	AttributeNameResponse struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
		// IsVerified bool   `json:"is_verified"`
	}
	AttributeNameUpdateRequest struct {
		Name string `json:"name" form:"name" binding:"omitempty,min=2,max=50"`
	}

	AttributeNameUpdateResponse struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
	AttributeValueCreateRequest struct {
		Name            string `json:"name" form:"name" validate:"required,min=2,max=50"`
		AttributeNameID string `json:"attribute_name" form:"attribute_name" validate:"required"`
	}
	AttributeNameWithValueResponse struct {
		ID     uint64                                       `json:"id"`
		Name   string                                       `json:"name"`
		Values helpers.PaginateData[AttributeValueResponse] `json:"values"`
	}
	AttributeValueResponse struct {
		ID              uint64  `json:"id"`
		Name            string  `json:"name" form:"name" `
		AttributeName   *string `json:"attribute_name,omitempty" `
		AttributeNameID *uint64 `json:"attribute_name_id,omitempty"`
	}

	AttributeValueUpdateRequest struct {
		Name            string `json:"name" form:"name" validate:"required,min=2,max=50"`
		AttributeNameID string `json:"attribute_name" form:"attribute_name" validate:"required"`
	}

	// AttributeNameUpdateResponse struct {
	// 	ID   uint64 `json:"id"`
	// 	Name string `json:"name"`
	// }
)
