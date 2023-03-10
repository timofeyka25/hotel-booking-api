package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"hotel-booking-app/internal/dao"
	"hotel-booking-app/internal/domain"
	"hotel-booking-app/pkg/customErrors"
	"hotel-booking-app/pkg/hash"
	"hotel-booking-app/pkg/jwt"
	"time"
)

type UserUseCase interface {
	SignUp(context.Context, SignUpParams) (uuid.UUID, error)
	SignIn(context.Context, SignInParams) (string, error)
	GetUser()
}

type userUseCase struct {
	dao            dao.UserDAO
	tokenGenerator *jwt.TokenGenerator
}

func NewUserUseCase(dao dao.UserDAO, t *jwt.TokenGenerator) *userUseCase {
	return &userUseCase{dao: dao, tokenGenerator: t}
}

func (uc userUseCase) SignUp(ctx context.Context, params SignUpParams) (uuid.UUID, error) {
	passwordHash, err := hash.ToHashString(params.Password)
	if err != nil {
		return uuid.Nil, errors.New(fmt.Sprintf("password hashing error: %s", err.Error()))
	}
	user := domain.NewUser(params.Name, params.Email, passwordHash, "be580d24-1eb1-4e5e-95d2-336e059167ae")

	if err = uc.dao.Create(ctx, user); err != nil {
		return uuid.Nil, err
	}

	return user.Id, nil
}

func (uc userUseCase) SignIn(ctx context.Context, params SignInParams) (string, error) {
	user, err := uc.dao.GetByEmail(ctx, params.Email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", customErrors.NewNotFoundError("user not found")
	}
	if !hash.IsEqualWithHash(params.Password, user.PasswordHash) {
		return "", errors.New("incorrect password")
	}

	return uc.tokenGenerator.GenerateNewAccessToken(24 * time.Hour)
}

func (uc userUseCase) GetUser() {
	//TODO implement me
	panic("implement me")
}

type SignUpParams struct {
	Name     string
	Email    string
	Password string
}

type SignInParams struct {
	Email    string
	Password string
}
