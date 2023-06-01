package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun/driver/pgdriver"
	"hotel-booking-app/internal/domain"
	"hotel-booking-app/pkg/custom_errors"
	"hotel-booking-app/pkg/db"
)

type RoleDAO interface {
	Create(ctx context.Context, role *domain.Role) error
	GetById(ctx context.Context, id uuid.UUID) (*domain.Role, error)
	GetByName(ctx context.Context, name string) (*domain.Role, error)
	Update(ctx context.Context, role *domain.Role) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type roleDAO struct {
	db *db.TransactionRepository
}

func NewRoleDAO(db *db.TransactionRepository) *roleDAO {
	return &roleDAO{db: db}
}

func (dao roleDAO) Create(ctx context.Context, role *domain.Role) error {
	_, err := dao.db.NewInsert(ctx).Model(role).Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		return custom_errors.NewAlreadyExistsError("role already exists")
	}

	return err
}

func (dao roleDAO) GetById(ctx context.Context, id uuid.UUID) (*domain.Role, error) {
	role := new(domain.Role)

	err := dao.db.NewSelect(ctx).
		Model(role).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (dao roleDAO) GetByName(ctx context.Context, name string) (*domain.Role, error) {
	role := new(domain.Role)

	err := dao.db.NewSelect(ctx).
		Model(role).
		Where("name = ?", name).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (dao roleDAO) Update(ctx context.Context, role *domain.Role) error {
	_, err := dao.db.NewUpdate(ctx).Model(role).Where("id = ?", role.Id).Exec(ctx)

	return err
}

func (dao roleDAO) Delete(ctx context.Context, id uuid.UUID) error {
	role := new(domain.Role)
	role.Id = id
	_, err := dao.db.NewDelete(ctx).Model(role).WherePK().Exec(ctx)

	return err
}
