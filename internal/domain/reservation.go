package domain

import (
	"github.com/google/uuid"
	"time"
)

const (
	PENDING     = "pending"
	CONFIRMED   = "confirmed"
	CANCELLED   = "cancelled"
	COMPLETED   = "completed"
	UNFULFILLED = "unfulfilled"
	PAID        = "paid"
	UNPAID      = "unpaid"
)

type Reservation struct {
	Id            uuid.UUID `bun:",pk"`
	UserId        uuid.UUID
	RoomId        uuid.UUID
	CheckInDate   time.Time
	CheckOutDate  time.Time
	Status        string
	PaymentStatus string
}

func NewReservation(
	userId, roomId string,
	CheckInDate, CheckOutDate time.Time,
	status, paymentStatus string) *Reservation {
	return &Reservation{
		Id:            uuid.New(),
		UserId:        uuid.MustParse(userId),
		RoomId:        uuid.MustParse(roomId),
		CheckInDate:   CheckInDate,
		CheckOutDate:  CheckOutDate,
		Status:        status,
		PaymentStatus: paymentStatus,
	}
}
