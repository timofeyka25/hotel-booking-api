package dao

import (
	"database/sql"
	"github.com/google/uuid"
	"hotel-booking-app/internal/domain"
)

type RoleDAO interface {
	Create(domain.Role) error
	Read(uuid.UUID) (*domain.Role, error)
	Update(domain.Role) error
	Delete(uuid.UUID) error
}

type roleDAO struct {
	db *sql.DB
}

func NewRoleDAO(db *sql.DB) *roleDAO {
	return &roleDAO{db: db}
}

func (dao roleDAO) Create(role domain.Role) error {
	_, err := dao.db.Exec("INSERT INTO roles (id, name) VALUES ($1, $2)", role.Id, role.Name)

	return err
}

func (dao roleDAO) Read(id uuid.UUID) (*domain.Role, error) {
	var r domain.Role
	err := dao.db.QueryRow("SELECT * FROM roles WHERE id = $1", id).Scan(&r.Id, &r.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &r, nil
}

func (dao roleDAO) Update(role domain.Role) error {
	_, err := dao.db.Exec("UPDATE roles SET name = $2 WHERE id = $1", role.Id, role.Name)

	return err
}

func (dao roleDAO) Delete(id uuid.UUID) error {
	_, err := dao.db.Exec("DELETE FROM roles WHERE id = $1", id)

	return err
}
