package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
	"hotel-booking-app/internal/domain"
	"hotel-booking-app/pkg/customErrors"
)

type RoomDAO interface {
	Create(context.Context, *domain.Room) error
	GetById(context.Context, uuid.UUID) (*domain.Room, error)
	GetByHotelId(context.Context, uuid.UUID) ([]*domain.Room, error)
	Update(context.Context, *domain.Room) error
	Delete(context.Context, uuid.UUID) error
}

type roomDAO struct {
	db *bun.DB
}

func NewRoomDAO(db *bun.DB) *roomDAO {
	return &roomDAO{db: db}
}

func (dao roomDAO) Create(ctx context.Context, room *domain.Room) error {
	_, err := dao.db.NewInsert().Model(room).Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		return customErrors.NewAlreadyExistsError("room already exists")
	}

	return err
}

func (dao roomDAO) GetById(ctx context.Context, id uuid.UUID) (*domain.Room, error) {
	room := new(domain.Room)

	err := dao.db.NewSelect().
		Model(room).
		Where("room.id = ?", id).
		Relation("Hotel").
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (dao roomDAO) GetByHotelId(ctx context.Context, hotelId uuid.UUID) ([]*domain.Room, error) {
	var rooms []*domain.Room
	err := dao.db.NewSelect().Model(&rooms).Where("hotel_id = ?", hotelId).Relation("Hotel").Scan(ctx)

	return rooms, err
}

func (dao roomDAO) Update(ctx context.Context, room *domain.Room) error {
	_, err := dao.db.NewUpdate().Model(room).Where("id = ?", room.Id).Exec(ctx)

	return err
}

func (dao roomDAO) Delete(ctx context.Context, id uuid.UUID) error {
	room := new(domain.Room)
	room.Id = id
	_, err := dao.db.NewDelete().Model(room).WherePK().Exec(ctx)

	return err
}
