package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
	"hotel-booking-app/internal/domain"
	"hotel-booking-app/pkg/customErrors"
)

type PaymentDAO interface {
	Create(context.Context, *domain.Payment) error
	GetById(context.Context, uuid.UUID) (*domain.Payment, error)
	Update(context.Context, *domain.Payment) error
	Delete(context.Context, uuid.UUID) error
}

type paymentDAO struct {
	db *bun.DB
}

func NewPaymentDAO(db *bun.DB) *paymentDAO {
	return &paymentDAO{db: db}
}

func (dao paymentDAO) Create(ctx context.Context, payment *domain.Payment) error {
	_, err := dao.db.NewInsert().Model(payment).Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		return customErrors.NewAlreadyExistsError("payment already exists")
	}

	return err
}

func (dao paymentDAO) GetById(ctx context.Context, id uuid.UUID) (*domain.Payment, error) {
	payment := new(domain.Payment)

	err := dao.db.NewSelect().
		Model(payment).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (dao paymentDAO) Update(ctx context.Context, payment *domain.Payment) error {
	_, err := dao.db.NewUpdate().Model(payment).Where("id = ?", payment.Id).Exec(ctx)

	return err
}

func (dao paymentDAO) Delete(ctx context.Context, id uuid.UUID) error {
	payment := new(domain.Payment)
	payment.Id = id
	_, err := dao.db.NewDelete().Model(payment).WherePK().Exec(ctx)

	return err
}
