package dto

import (
	"github.com/google/uuid"
	"hotel-booking-app/pkg/custom_errors"
	"time"
)

const layout = "2006-01-02"

type CreateReservationDTO struct {
	CheckInDate  string `json:"check_in_date" validate:"required,datetime=2006-01-02"`
	CheckOutDate string `json:"check_out_date" validate:"required,datetime=2006-01-02"`
}

type CreateReservationParsedDTO struct {
	CheckInDate  time.Time
	CheckOutDate time.Time
}

func (c *CreateReservationDTO) ParseAndValidate() (*CreateReservationParsedDTO, error) {
	checkInDate, err := time.Parse(layout, c.CheckInDate)
	if err != nil {
		return nil, err
	}
	checkOutDate, err := time.Parse(layout, c.CheckOutDate)
	if err != nil {
		return nil, err
	}
	today := time.Now()
	if checkInDate.Before(today) || checkOutDate.Before(today) || checkOutDate.Before(checkInDate) {
		return nil, custom_errors.NewValidationError(
			"CheckInDate and CheckOutDate must be greater than today and CheckOutDate must be greater than CheckInDate")
	}
	return &CreateReservationParsedDTO{CheckInDate: checkInDate, CheckOutDate: checkOutDate}, nil
}

type ReservationDTO struct {
	Id            uuid.UUID `json:"id"`
	UserId        uuid.UUID `json:"user_id"`
	Room          *RoomDTO  `json:"room"`
	CheckInDate   time.Time `json:"check_in_date"`
	CheckOutDate  time.Time `json:"check_out_date"`
	Status        string    `json:"status"`
	PaymentStatus string    `json:"payment_status"`
}

type UpdateReservationStatusDTO struct {
	Status string `json:"status"`
}
