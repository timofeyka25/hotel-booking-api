package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun/driver/pgdriver"
	"hotel-booking-app/internal/domain"
	"hotel-booking-app/pkg/custom_errors"
	"hotel-booking-app/pkg/db"
)

type ReservationDAO interface {
	Create(ctx context.Context, reservation *domain.Reservation) error
	GetById(ctx context.Context, id uuid.UUID) (*domain.Reservation, error)
	GetByRoomAndUserId(ctx context.Context, userID, roomId uuid.UUID) ([]*domain.Reservation, error)
	Update(ctx context.Context, reservation *domain.Reservation) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByUserId(ctx context.Context, userId uuid.UUID) ([]*domain.Reservation, error)
	GetAll(ctx context.Context) ([]*domain.Reservation, error)
}

type reservationDAO struct {
	db *db.TransactionRepository
}

func NewReservationDAO(db *db.TransactionRepository) *reservationDAO {
	return &reservationDAO{db: db}
}

func (dao reservationDAO) Create(ctx context.Context, r *domain.Reservation) error {
	_, err := dao.db.NewInsert(ctx).Model(r).Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		return custom_errors.NewAlreadyExistsError("reservation already exists")
	}

	return err
}

func (dao reservationDAO) GetById(ctx context.Context, id uuid.UUID) (*domain.Reservation, error) {
	r := new(domain.Reservation)

	err := dao.db.NewSelect(ctx).
		Model(r).
		Where("reservation.id = ?", id).
		Relation("Room").
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return r, nil
}
func (dao reservationDAO) GetByRoomAndUserId(
	ctx context.Context,
	userId, roomId uuid.UUID,
) ([]*domain.Reservation, error) {
	var r []*domain.Reservation

	err := dao.db.NewSelect(ctx).
		Model(&r).
		Where("user_id = ?", userId).
		Where("room_id = ?", roomId).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (dao reservationDAO) GetByUserId(ctx context.Context, id uuid.UUID) ([]*domain.Reservation, error) {
	var r []*domain.Reservation

	err := dao.db.NewSelect(ctx).
		Model(&r).
		Where("reservation.user_id = ?", id).
		Relation("Room").
		Relation("Room.Hotel").
		OrderExpr("reservation.check_in_date ASC").Scan(ctx)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (dao reservationDAO) GetAll(ctx context.Context) ([]*domain.Reservation, error) {
	var r []*domain.Reservation
	err := dao.db.NewSelect(ctx).
		Model(&r).
		Relation("Room").
		Relation("Room.Hotel").
		OrderExpr("reservation.check_in_date ASC").Scan(ctx)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (dao reservationDAO) Update(ctx context.Context, r *domain.Reservation) error {
	_, err := dao.db.NewUpdate(ctx).Model(r).Where("id = ?", r.Id).Exec(ctx)

	return err
}

func (dao reservationDAO) Delete(ctx context.Context, id uuid.UUID) error {
	r := new(domain.Reservation)
	r.Id = id
	_, err := dao.db.NewDelete(ctx).Model(r).WherePK().Exec(ctx)

	return err
}
