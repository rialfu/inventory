package dto

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
	MerkCreateRequest struct {
		Name string `json:"name" form:"name" validate:"required,min=2,max=50"`
	}

	MerkResponse struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
		// IsVerified bool   `json:"is_verified"`
	}
	MerkUpdateRequest struct {
		Name string `json:"name" form:"name" validate:"required,min=2,max=50"`
	}

	MerkUpdateResponse struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
	MerkDropdownDTO struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
)
