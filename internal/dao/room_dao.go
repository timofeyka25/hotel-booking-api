package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun/driver/pgdriver"
	"hotel-booking-app/internal/domain"
	"hotel-booking-app/pkg/custom_errors"
	"hotel-booking-app/pkg/db"
	"time"
)

type RoomDAO interface {
	Create(ctx context.Context, room *domain.Room) error
	GetById(ctx context.Context, id uuid.UUID) (*domain.Room, error)
	GetByHotelId(ctx context.Context, hotelId uuid.UUID) ([]*domain.Room, error)
	GetByHotelIdFreeRooms(ctx context.Context, hotelId uuid.UUID) ([]*domain.Room, error)
	Update(ctx context.Context, room *domain.Room) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type roomDAO struct {
	db *db.TransactionRepository
}

func NewRoomDAO(db *db.TransactionRepository) *roomDAO {
	return &roomDAO{db: db}
}

func (dao roomDAO) Create(ctx context.Context, room *domain.Room) error {
	_, err := dao.db.NewInsert(ctx).Model(room).Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		return custom_errors.NewAlreadyExistsError("room already exists")
	}

	return err
}

func (dao roomDAO) GetById(ctx context.Context, id uuid.UUID) (*domain.Room, error) {
	room := new(domain.Room)

	err := dao.db.NewSelect(ctx).
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
	err := dao.db.NewSelect(ctx).Model(&rooms).Where("hotel_id = ?", hotelId).Relation("Hotel").Scan(ctx)

	return rooms, err
}

func (dao roomDAO) GetByHotelIdFreeRooms(ctx context.Context, hotelId uuid.UUID) ([]*domain.Room, error) {
	var rooms []*domain.Room
	err := dao.db.NewSelect(ctx).
		Model(&rooms).
		Where("hotel_id = ?", hotelId).
		Relation("Hotel").
		Join("LEFT JOIN reservations ON room.id = reservations.room_id").
		Where("reservations.check_out_date IS NULL").
		WhereOr("reservations.check_out_date < ?", time.Now()).
		WhereOr("reservations.status IS NULL").
		WhereOr("reservations.status = ?", domain.CANCELLED).
		Scan(ctx)

	return rooms, err
}

func (dao roomDAO) Update(ctx context.Context, room *domain.Room) error {
	_, err := dao.db.NewUpdate(ctx).Model(room).Where("id = ?", room.Id).Exec(ctx)

	return err
}

func (dao roomDAO) Delete(ctx context.Context, id uuid.UUID) error {
	room := new(domain.Room)
	room.Id = id
	_, err := dao.db.NewDelete(ctx).Model(room).WherePK().Exec(ctx)

	return err
}
