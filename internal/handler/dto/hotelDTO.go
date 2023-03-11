package dto

import "github.com/google/uuid"

type AddHotelReqDTO struct {
	Name        string `json:"name" validate:"required"`
	Location    string `json:"location" validate:"required"`
	Description string `json:"description"`
}

type AddHotelResDTO struct {
	Id      uuid.UUID `json:"id,omitempty"`
	Message string    `json:"message,omitempty"`
}
