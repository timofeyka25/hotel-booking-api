package dao

import (
	"database/sql"
	"github.com/google/uuid"
	"hotel-booking-app/domain"
)

type UserDAO interface {
	Create(domain.User) error
	Read(uuid.UUID) (*domain.User, error)
	Update(domain.User) error
	Delete(uuid.UUID) error
}

type userDAO struct {
	db *sql.DB
}

func NewUserDAO(db *sql.DB) *userDAO {
	return &userDAO{db: db}
}

func (dao userDAO) Create(user domain.User) error {
	_, err := dao.db.Exec(`INSERT INTO users (id, name, email, password_hash, role_id) VALUES ($1, $2, $3, $4, $5)`,
		user.Id, user.Name, user.Email, user.Password, user.RoleId)

	return err
}

func (dao userDAO) Read(id uuid.UUID) (*domain.User, error) {
	var user domain.User
	err := dao.db.QueryRow("SELECT id, name, email, password_hash, role_id, is_active FROM users WHERE id = $1", id).
		Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.RoleId, &user.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no rows found
		}
		return nil, err // some other error occurred
	}
	return &user, nil // user found
}

func (dao userDAO) Update(user domain.User) error {
	_, err := dao.db.Exec(`UPDATE users SET name = $2, email = $3, password_hash = $4, role_id = $5, is_active = $6 WHERE id = $1`,
		user.Id, user.Name, user.Email, user.Password, user.RoleId, user.IsActive)

	return err
}

func (dao userDAO) Delete(id uuid.UUID) error {
	_, err := dao.db.Exec(`DELETE FROM users WHERE id = $1`, id)

	return err
}
