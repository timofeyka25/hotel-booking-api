package dto

import (
	"github.com/google/uuid"
	"time"
)

type CreatePaymentDTO struct {
	Amount float64 `json:"amount" validate:"required"`
}

type PaymentDTO struct {
	Id            uuid.UUID `json:"id"`
	ReservationId uuid.UUID `json:"reservation_id"`
	UserId        uuid.UUID `json:"user_id"`
	Amount        float64   `json:"amount"`
	PaymentTime   time.Time `json:"payment_time"`
}
