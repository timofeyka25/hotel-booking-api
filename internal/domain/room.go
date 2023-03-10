package domain

import "github.com/google/uuid"

type Room struct {
	Id            uuid.UUID `bun:",pk"`
	HotelId       uuid.UUID
	RoomType      string
	MaxOccupancy  int
	PricePerNight float64
}

func NewRoom(
	hotelId uuid.UUID,
	roomType string,
	maxOccupancy int,
	pricePerNight float64) Room {
	return Room{
		Id:            uuid.New(),
		HotelId:       hotelId,
		RoomType:      roomType,
		MaxOccupancy:  maxOccupancy,
		PricePerNight: pricePerNight,
	}
}
