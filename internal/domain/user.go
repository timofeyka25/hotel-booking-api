package domain

import "github.com/google/uuid"

type User struct {
	Id           uuid.UUID `bun:",pk"`
	Name         string
	Email        string
	PasswordHash string `bun:"password_hash"`
	RoleId       uuid.UUID
	IsActive     bool
}

func NewUser(
	name string,
	email string,
	password string,
	roleId string) *User {
	return &User{
		Id:           uuid.New(),
		Name:         name,
		Email:        email,
		PasswordHash: password,
		RoleId:       uuid.MustParse(roleId),
		IsActive:     true,
	}
}
