package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun/driver/pgdriver"
	"hotel-booking-app/internal/domain"
	"hotel-booking-app/pkg/customErrors"
	"hotel-booking-app/pkg/db"
)

type PaymentDAO interface {
	Create(ctx context.Context, payment *domain.Payment) error
	GetById(ctx context.Context, id uuid.UUID) (*domain.Payment, error)
	GetByUserId(ctx context.Context, userId uuid.UUID) ([]*domain.Payment, error)
	Update(ctx context.Context, payment *domain.Payment) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type paymentDAO struct {
	db *db.TransactionRepository
}

func NewPaymentDAO(db *db.TransactionRepository) *paymentDAO {
	return &paymentDAO{db: db}
}

func (dao paymentDAO) Create(ctx context.Context, payment *domain.Payment) error {
	_, err := dao.db.NewInsert(ctx).Model(payment).Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		return customErrors.NewAlreadyExistsError("payment already exists")
	}

	return err
}

func (dao paymentDAO) GetById(ctx context.Context, id uuid.UUID) (*domain.Payment, error) {
	payment := new(domain.Payment)

	err := dao.db.NewSelect(ctx).
		Model(payment).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (dao paymentDAO) GetByUserId(ctx context.Context, userId uuid.UUID) ([]*domain.Payment, error) {
	var payments []*domain.Payment

	err := dao.db.NewSelect(ctx).
		Model(&payments).
		Where("user_id = ?", userId).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (dao paymentDAO) Update(ctx context.Context, payment *domain.Payment) error {
	_, err := dao.db.NewUpdate(ctx).Model(payment).Where("id = ?", payment.Id).Exec(ctx)

	return err
}

func (dao paymentDAO) Delete(ctx context.Context, id uuid.UUID) error {
	payment := new(domain.Payment)
	payment.Id = id
	_, err := dao.db.NewDelete(ctx).Model(payment).WherePK().Exec(ctx)

	return err
}
