package domain

import (
	"github.com/google/uuid"
	"time"
)

type Payment struct {
	Id            uuid.UUID
	ReservationId uuid.UUID
	UserId        uuid.UUID
	Amount        float64
	PaymentTime   time.Time
}

func NewPayment(
	reservationId, userId uuid.UUID,
	amount float64) Payment {
	return Payment{
		Id:            uuid.New(),
		ReservationId: reservationId,
		UserId:        userId,
		Amount:        amount,
		PaymentTime:   time.Now(),
	}
}
