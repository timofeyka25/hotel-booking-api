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
	GetUser(context.Context, string) (*domain.User, error)
}

type userUseCase struct {
	userDAO        dao.UserDAO
	roleDAO        dao.RoleDAO
	tokenGenerator *jwt.TokenGenerator
}

func NewUserUseCase(userDAO dao.UserDAO, roleDAO dao.RoleDAO, t *jwt.TokenGenerator) *userUseCase {
	return &userUseCase{userDAO: userDAO, roleDAO: roleDAO, tokenGenerator: t}
}

func (uc userUseCase) SignUp(ctx context.Context, params SignUpParams) (uuid.UUID, error) {
	passwordHash, err := hash.ToHashString(params.Password)
	if err != nil {
		return uuid.Nil, errors.New(fmt.Sprintf("password hashing error: %s", err.Error()))
	}
	role, err := uc.roleDAO.GetByName(ctx, domain.USER)
	if err != nil {
		return uuid.Nil, err
	}
	user := domain.NewUser(params.Name, params.Email, passwordHash, role.Id)

	if err = uc.userDAO.Create(ctx, user); err != nil {
		return uuid.Nil, err
	}

	return user.Id, nil
}

func (uc userUseCase) SignIn(ctx context.Context, params SignInParams) (string, error) {
	user, err := uc.userDAO.GetByEmail(ctx, params.Email)
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

func (uc userUseCase) GetUser(ctx context.Context, email string) (*domain.User, error) {
	user, err := uc.userDAO.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, customErrors.NewNotFoundError("user not found")
	}

	return user, nil
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
