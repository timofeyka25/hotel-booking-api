package dto

import "github.com/google/uuid"

type AddHotelReqDTO struct {
	Name        string `json:"name" validate:"required"`
	Location    string `json:"location" validate:"required"`
	Description string `json:"description"`
}

type AddHotelResDTO struct {
	Id uuid.UUID `json:"id,omitempty"`
}

type HotelDTO struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
}

type AllHotelsDTO struct {
	Hotels []HotelDTO `json:"hotels,omitempty"`
}

type UpdateHotelDTO struct {
	Name        *string `json:"name,omitempty"`
	Location    *string `json:"location,omitempty"`
	Description *string `json:"description,omitempty"`
}
