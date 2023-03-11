package usecase

import (
	"context"
	"github.com/google/uuid"
	"hotel-booking-app/internal/dao"
	"hotel-booking-app/internal/domain"
)

type HotelUseCase interface {
	AddHotel(context.Context, AddHotelParams) (uuid.UUID, error)
}

type hotelUseCase struct {
	hotelDAO dao.HotelDAO
}

func NewHotelUseCase(hotelDAO dao.HotelDAO) *hotelUseCase {
	return &hotelUseCase{hotelDAO: hotelDAO}
}

func (h hotelUseCase) AddHotel(ctx context.Context, params AddHotelParams) (uuid.UUID, error) {
	hotel := domain.NewHotel(params.Name, params.Location, params.Description)
	if err := h.hotelDAO.Create(ctx, hotel); err != nil {
		return uuid.Nil, err
	}
	return hotel.Id, nil
}

type AddHotelParams struct {
	Name        string
	Location    string
	Description string
}
