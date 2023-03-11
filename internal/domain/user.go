package domain

import "github.com/google/uuid"

type User struct {
	Id           uuid.UUID `bun:",pk"`
	Name         string
	Email        string
	PasswordHash string `bun:"password_hash"`
	IsActive     bool
	RoleId       uuid.UUID
	Role         *Role `bun:"rel:belongs-to"`
}

func NewUser(
	name string,
	email string,
	password string,
	roleId uuid.UUID) *User {
	return &User{
		Id:           uuid.New(),
		Name:         name,
		Email:        email,
		PasswordHash: password,
		RoleId:       roleId,
		IsActive:     true,
	}
}
