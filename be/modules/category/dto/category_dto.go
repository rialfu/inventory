package dto

import "errors"

var (
	ErrGetParentNotFound = errors.New("parent not found")
	ErrDataCircular      = errors.New("data circular")
)

type (
	CategoryCreateRequest struct {
		Name     string `json:"name" form:"name" validate:"required,min=2,max=100"`
		ParentID string `json:"parent_id"`
	}

	CategoryResponse struct {
		ID       uint64  `json:"id"`
		Name     string  `json:"name"`
		Path     *string `json:"path"`
		ParentID *uint64 `json:"parent_id"`
	}

	CategoryDropdownDTO struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}

	CategoryUpdateRequest struct {
		Name     string `json:"name" form:"name" validate:"required,min=2,max=100"`
		ParentID string `json:"parent_id"`
	}
)
