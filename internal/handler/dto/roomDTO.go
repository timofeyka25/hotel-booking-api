package dto

import "github.com/google/uuid"

type AddRoomDTO struct {
	RoomType      string  `json:"room_type"`
	MaxOccupancy  int     `json:"max_occupancy"`
	PricePerNight float64 `json:"price_per_night"`
}

type RoomDTO struct {
	Id            uuid.UUID `json:"id"`
	Hotel         *HotelDTO `json:"hotel"`
	RoomType      string    `json:"room_type"`
	MaxOccupancy  int       `json:"max_occupancy"`
	PricePerNight float64   `json:"price_per_night"`
}
