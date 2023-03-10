package dto

import "github.com/google/uuid"

type SignInRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=255"`
}

type SignInResponseDTO struct {
	Token   string `json:"token,omitempty"`
	Message string `json:"message,omitempty"`
}

type SignUpRequestDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=255"`
}

type SignUpResponseDTO struct {
	Id      uuid.UUID `json:"id,omitempty"`
	Message string    `json:"message,omitempty"`
}
