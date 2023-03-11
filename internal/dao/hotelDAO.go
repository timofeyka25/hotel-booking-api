package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
	"hotel-booking-app/internal/domain"
	"hotel-booking-app/pkg/customErrors"
)

type HotelDAO interface {
	Create(context.Context, *domain.Hotel) error
	GetById(context.Context, uuid.UUID) (*domain.Hotel, error)
	Update(context.Context, *domain.Hotel) error
	Delete(context.Context, uuid.UUID) error
}

type hotelDAO struct {
	db *bun.DB
}

func NewHotelDAO(db *bun.DB) *hotelDAO {
	return &hotelDAO{db: db}
}

func (dao hotelDAO) Create(ctx context.Context, hotel *domain.Hotel) error {
	_, err := dao.db.NewInsert().Model(hotel).Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		return customErrors.NewAlreadyExistsError("hotel already exists")
	}

	return err
}

func (dao hotelDAO) GetById(ctx context.Context, id uuid.UUID) (*domain.Hotel, error) {
	hotel := new(domain.Hotel)

	err := dao.db.NewSelect().
		Model(hotel).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return hotel, nil
}

func (dao hotelDAO) Update(ctx context.Context, hotel *domain.Hotel) error {
	_, err := dao.db.NewUpdate().Model(hotel).Where("id = ?", hotel.Id).Exec(ctx)

	return err
}

func (dao hotelDAO) Delete(ctx context.Context, id uuid.UUID) error {
	hotel := new(domain.Hotel)
	hotel.Id = id
	_, err := dao.db.NewDelete().Model(hotel).WherePK().Exec(ctx)

	return err
}
