package dao

import (
	"database/sql"
	"github.com/google/uuid"
	"hotel-booking-app/internal/domain"
)

type PaymentDAO interface {
	Create(domain.Payment) error
	Read(uuid.UUID) (*domain.Payment, error)
	Update(domain.Payment) error
	Delete(uuid.UUID) error
}

type paymentDAO struct {
	db *sql.DB
}

func NewPaymentDAO(db *sql.DB) *paymentDAO {
	return &paymentDAO{db: db}
}

func (dao paymentDAO) Create(p domain.Payment) error {
	//TODO implement me
	panic("implement me")
}

func (dao paymentDAO) Read(id uuid.UUID) (*domain.Payment, error) {
	//TODO implement me
	panic("implement me")
}

func (dao paymentDAO) Update(p domain.Payment) error {
	//TODO implement me
	panic("implement me")
}

func (dao paymentDAO) Delete(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
