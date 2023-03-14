package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun/driver/pgdriver"
	"hotel-booking-app/internal/domain"
	"hotel-booking-app/pkg/customErrors"
	"hotel-booking-app/pkg/db"
)

type HotelDAO interface {
	Create(ctx context.Context, hotel *domain.Hotel) error
	GetById(ctx context.Context, id uuid.UUID) (*domain.Hotel, error)
	GetAll(ctx context.Context) ([]*domain.Hotel, error)
	Update(ctx context.Context, hotel *domain.Hotel) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type hotelDAO struct {
	db *db.TransactionRepository
}

func NewHotelDAO(db *db.TransactionRepository) *hotelDAO {
	return &hotelDAO{db: db}
}

func (dao hotelDAO) Create(ctx context.Context, hotel *domain.Hotel) error {
	_, err := dao.db.NewInsert(ctx).Model(hotel).Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		return customErrors.NewAlreadyExistsError("hotel already exists")
	}

	return err
}

func (dao hotelDAO) GetById(ctx context.Context, id uuid.UUID) (*domain.Hotel, error) {
	hotel := new(domain.Hotel)

	err := dao.db.NewSelect(ctx).
		Model(hotel).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return hotel, nil
}

func (dao hotelDAO) GetAll(ctx context.Context) ([]*domain.Hotel, error) {
	var hotels []*domain.Hotel
	err := dao.db.NewSelect(ctx).Model(&hotels).Scan(ctx)

	return hotels, err
}

func (dao hotelDAO) Update(ctx context.Context, hotel *domain.Hotel) error {
	_, err := dao.db.NewUpdate(ctx).Model(hotel).Where("id = ?", hotel.Id).Exec(ctx)

	return err
}

func (dao hotelDAO) Delete(ctx context.Context, id uuid.UUID) error {
	hotel := new(domain.Hotel)
	hotel.Id = id
	_, err := dao.db.NewDelete(ctx).Model(hotel).WherePK().Exec(ctx)

	return err
}
