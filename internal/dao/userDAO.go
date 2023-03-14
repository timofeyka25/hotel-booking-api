package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun/driver/pgdriver"
	"hotel-booking-app/internal/domain"
	"hotel-booking-app/pkg/customErrors"
	"hotel-booking-app/pkg/db"
)

type UserDAO interface {
	Create(ctx context.Context, user *domain.User) error
	GetById(ctx context.Context, id uuid.UUID) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetUsers(ctx context.Context) ([]*domain.User, error)
}

type userDAO struct {
	db *db.TransactionRepository
}

func NewUserDAO(db *db.TransactionRepository) *userDAO {
	return &userDAO{db: db}
}

func (dao userDAO) Create(ctx context.Context, user *domain.User) error {
	_, err := dao.db.NewInsert(ctx).Model(user).Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		return customErrors.NewAlreadyExistsError("user already exists")
	}

	return err
}

func (dao userDAO) GetById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user := new(domain.User)

	err := dao.db.NewSelect(ctx).
		Model(user).
		Where("\"user\".\"id\" = ?", id).
		Relation("Role").
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (dao userDAO) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := new(domain.User)

	err := dao.db.NewSelect(ctx).
		Model(user).
		Where("email = ?", email).
		Relation("Role").
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (dao userDAO) GetUsers(ctx context.Context) ([]*domain.User, error) {
	var users []*domain.User

	err := dao.db.NewSelect(ctx).
		Column("id", "name", "email", "is_active", "role_id").
		Model(&users).
		Relation("Role").
		Scan(ctx)

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (dao userDAO) Update(ctx context.Context, user *domain.User) error {
	_, err := dao.db.NewUpdate(ctx).Model(user).Where("id = ?", user.Id).Exec(ctx)

	return err
}

func (dao userDAO) Delete(ctx context.Context, id uuid.UUID) error {
	u := new(domain.User)
	u.Id = id
	_, err := dao.db.NewDelete(ctx).Model(u).WherePK().Exec(ctx)

	return err
}
