package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
	"hotel-booking-app/internal/domain"
	"hotel-booking-app/pkg/customErrors"
)

type RoleDAO interface {
	Create(context.Context, *domain.Role) error
	GetById(context.Context, uuid.UUID) (*domain.Role, error)
	GetByName(context.Context, string) (*domain.Role, error)
	Update(context.Context, *domain.Role) error
	Delete(context.Context, uuid.UUID) error
}

type roleDAO struct {
	db *bun.DB
}

func NewRoleDAO(db *bun.DB) *roleDAO {
	return &roleDAO{db: db}
}

func (dao roleDAO) Create(ctx context.Context, role *domain.Role) error {
	_, err := dao.db.NewInsert().Model(role).Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		return customErrors.NewAlreadyExistsError("role already exists")
	}

	return err
}

func (dao roleDAO) GetById(ctx context.Context, id uuid.UUID) (*domain.Role, error) {
	role := new(domain.Role)

	err := dao.db.NewSelect().
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

	err := dao.db.NewSelect().
		Model(role).
		Where("name = ?", name).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (dao roleDAO) Update(ctx context.Context, role *domain.Role) error {
	_, err := dao.db.NewUpdate().Model(role).Where("id = ?", role.Id).Exec(ctx)

	return err
}

func (dao roleDAO) Delete(ctx context.Context, id uuid.UUID) error {
	role := new(domain.Role)
	role.Id = id
	_, err := dao.db.NewDelete().Model(role).WherePK().Exec(ctx)

	return err
}
