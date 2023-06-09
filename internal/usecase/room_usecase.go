package usecase

import (
	"context"
	"github.com/google/uuid"
	"hotel-booking-app/internal/dao"
	"hotel-booking-app/internal/domain"
)

type RoomUseCase interface {
	AddRoom(ctx context.Context, params AddRoomParams) (uuid.UUID, error)
	GetRoomById(ctx context.Context, id uuid.UUID) (*domain.Room, error)
	GetHotelRooms(ctx context.Context, hotelId uuid.UUID) ([]*domain.Room, error)
	GetHotelFreeRooms(ctx context.Context, hotelId uuid.UUID) ([]*domain.Room, error)
	UpdateRoom(ctx context.Context, params UpdateRoomParams) error
	DeleteRoom(ctx context.Context, id uuid.UUID) error
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

func (uc roomUseCase) UpdateRoom(ctx context.Context, params UpdateRoomParams) error {
	room, err := uc.roomDAO.GetById(ctx, params.Id)
	if err != nil {
		return err
	}
	if params.RoomType != nil {
		room.RoomType = *params.RoomType
	}
	if params.MaxOccupancy != nil {
		room.MaxOccupancy = *params.MaxOccupancy
	}
	if params.PricePerNight != nil {
		room.PricePerNight = *params.PricePerNight
	}
	return uc.roomDAO.Update(ctx, room)
}

func (uc roomUseCase) DeleteRoom(ctx context.Context, id uuid.UUID) error {
	return uc.roomDAO.Delete(ctx, id)
}

type AddRoomParams struct {
	HotelId       uuid.UUID
	RoomType      string
	MaxOccupancy  int
	PricePerNight float64
}

type UpdateRoomParams struct {
	Id            uuid.UUID
	RoomType      *string
	MaxOccupancy  *int
	PricePerNight *float64
}
