package usecase

import (
	"context"
	"github.com/google/uuid"
	"hotel-booking-app/internal/dao"
	"hotel-booking-app/internal/domain"
)

type RoomUseCase interface {
	AddRoom(context.Context, AddRoomParams) (uuid.UUID, error)
	GetRoomById(context.Context, uuid.UUID) (*domain.Room, error)
	GetHotelRooms(context.Context, uuid.UUID) ([]*domain.Room, error)
	GetHotelFreeRooms(context.Context, uuid.UUID) ([]*domain.Room, error)
}

type roomUseCase struct {
	roomDAO dao.RoomDAO
}

func NewRoomUseCase(roomDAO dao.RoomDAO) *roomUseCase {
	return &roomUseCase{roomDAO: roomDAO}
}

func (uc roomUseCase) AddRoom(ctx context.Context, params AddRoomParams) (uuid.UUID, error) {
	room := domain.NewRoom(
		params.HotelId,
		params.RoomType,
		params.MaxOccupancy,
		params.PricePerNight,
	)
	if err := uc.roomDAO.Create(ctx, room); err != nil {
		return uuid.Nil, err
	}
	return room.Id, nil
}

func (uc roomUseCase) GetRoomById(ctx context.Context, id uuid.UUID) (*domain.Room, error) {
	return uc.roomDAO.GetById(ctx, id)
}

func (uc roomUseCase) GetHotelRooms(ctx context.Context, hotelId uuid.UUID) ([]*domain.Room, error) {
	return uc.roomDAO.GetByHotelId(ctx, hotelId)
}

func (uc roomUseCase) GetHotelFreeRooms(ctx context.Context, hotelId uuid.UUID) ([]*domain.Room, error) {
	return uc.roomDAO.GetByHotelIdFreeRooms(ctx, hotelId)
}

type AddRoomParams struct {
	HotelId       uuid.UUID
	RoomType      string
	MaxOccupancy  int
	PricePerNight float64
}
