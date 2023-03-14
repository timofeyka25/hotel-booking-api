package usecase

import (
	"context"
	"github.com/google/uuid"
	"hotel-booking-app/internal/dao"
	"hotel-booking-app/internal/domain"
)

type HotelUseCase interface {
	AddHotel(ctx context.Context, params AddHotelParams) (uuid.UUID, error)
	GetAllHotels(ctx context.Context) ([]*domain.Hotel, error)
	GetHotelById(ctx context.Context, id uuid.UUID) (*domain.Hotel, error)
	UpdateHotel(ctx context.Context, params UpdateHotelParams) error
	DeleteHotel(ctx context.Context, id uuid.UUID) error
}

type hotelUseCase struct {
	hotelDAO dao.HotelDAO
}

func NewHotelUseCase(hotelDAO dao.HotelDAO) *hotelUseCase {
	return &hotelUseCase{hotelDAO: hotelDAO}
}

func (uc hotelUseCase) AddHotel(ctx context.Context, params AddHotelParams) (uuid.UUID, error) {
	hotel := domain.NewHotel(params.Name, params.Location, params.Description)
	if err := uc.hotelDAO.Create(ctx, hotel); err != nil {
		return uuid.Nil, err
	}
	return hotel.Id, nil
}

func (uc hotelUseCase) GetAllHotels(ctx context.Context) ([]*domain.Hotel, error) {
	return uc.hotelDAO.GetAll(ctx)
}

func (uc hotelUseCase) GetHotelById(ctx context.Context, id uuid.UUID) (*domain.Hotel, error) {
	return uc.hotelDAO.GetById(ctx, id)
}

func (uc hotelUseCase) UpdateHotel(ctx context.Context, params UpdateHotelParams) error {
	hotel, err := uc.hotelDAO.GetById(ctx, params.Id)
	if err != nil {
		return err
	}
	if params.Name != nil {
		hotel.Name = *params.Name
	}
	if params.Location != nil {
		hotel.Location = *params.Location
	}
	if params.Description != nil {
		hotel.Description = *params.Description
	}
	return uc.hotelDAO.Update(ctx, hotel)
}

func (uc hotelUseCase) DeleteHotel(ctx context.Context, id uuid.UUID) error {
	return uc.hotelDAO.Delete(ctx, id)
}

type AddHotelParams struct {
	Name        string
	Location    string
	Description string
}

type UpdateHotelParams struct {
	Id          uuid.UUID
	Name        *string
	Location    *string
	Description *string
}
