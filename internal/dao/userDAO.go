package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
	"hotel-booking-app/internal/domain"
	"hotel-booking-app/pkg/customErrors"
)

type UserDAO interface {
	Create(context.Context, *domain.User) error
	GetById(context.Context, uuid.UUID) (*domain.User, error)
	GetByEmail(context.Context, string) (*domain.User, error)
	Update(context.Context, *domain.User) error
	Delete(context.Context, uuid.UUID) error
}

type userDAO struct {
	db *bun.DB
}

func NewUserDAO(db *bun.DB) *userDAO {
	return &userDAO{db: db}
}

func (dao userDAO) Create(ctx context.Context, user *domain.User) error {
	_, err := dao.db.NewInsert().Model(user).Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		return customErrors.NewAlreadyExistsError("user already exists")
	}

	return err
}

func (dao userDAO) GetById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user := new(domain.User)

	err := dao.db.NewSelect().
		Model(user).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (dao userDAO) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := new(domain.User)

	err := dao.db.NewSelect().
		Model(user).
		Where("email = ?", email).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (dao userDAO) Update(ctx context.Context, user *domain.User) error {
	_, err := dao.db.NewUpdate().Model(user).Where("id = ?", user.Id).Exec(ctx)

	return err
}

func (dao userDAO) Delete(ctx context.Context, id uuid.UUID) error {
	u := new(domain.User)
	u.Id = id
	_, err := dao.db.NewDelete().Model(u).WherePK().Exec(ctx)

	return err
}
