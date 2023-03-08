package domain

import (
	"github.com/google/uuid"
	"time"
)

type Reservation struct {
	Id           uuid.UUID
	UserId       uuid.UUID
	RoomId       uuid.UUID
	CheckInDate  time.Time
	CheckOutDate time.Time
	Status       string
}

func NewReservation(
	userId, roomId string,
	CheckInDate, CheckOutDate time.Time,
	status string) Reservation {
	return Reservation{
		Id:           uuid.New(),
		UserId:       uuid.MustParse(userId),
		RoomId:       uuid.MustParse(roomId),
		CheckInDate:  CheckInDate,
		CheckOutDate: CheckOutDate,
		Status:       status,
	}
}
