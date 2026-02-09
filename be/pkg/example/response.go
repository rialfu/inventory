package example

import (
	authDTO "rialfu/wallet/modules/auth/dto"
	userDTO "rialfu/wallet/modules/user/dto"
)

type ResponseRegister struct {
	Status  bool                 `json:"status"`
	Message string               `json:"message"`
	Data    userDTO.UserResponse `json:"data,omitempty"`
}

type ResponseLogin struct {
	Status  bool                  `json:"status"`
	Message string                `json:"message"`
	Data    authDTO.TokenResponse `json:"data,omitempty"`
}

