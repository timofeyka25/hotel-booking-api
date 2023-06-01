package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"hotel-booking-app/internal/dao"
	"hotel-booking-app/internal/domain"
	"hotel-booking-app/pkg/custom_errors"
	"hotel-booking-app/pkg/hash"
	"hotel-booking-app/pkg/jwt"
	"time"
)

type UserUseCase interface {
	SignUp(ctx context.Context, params SignUpParams) (uuid.UUID, error)
	SignIn(ctx context.Context, params SignInParams) (string, error)
	GetUser(ctx context.Context, email string) (*domain.User, error)
	GetUsersList(ctx context.Context) ([]*domain.User, error)
	UpdateUserActiveStatus(ctx context.Context, params UpdateUserActiveParams) error
	UpdateUserRole(ctx context.Context, params UpdateUserRoleParams) error
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
		return "", custom_errors.NewNotFoundError("user not found")
	}
	if !hash.IsEqualWithHash(params.Password, user.PasswordHash) {
		return "", errors.New("incorrect password")
	}

	if user.IsActive == false {
		return "", custom_errors.NewNotActiveError("This account is no longer active")
	}

	return uc.tokenGenerator.GenerateNewAccessToken(jwt.Params{
		Id:   user.Id.String(),
		Role: user.Role.Name,
		Ttl:  24 * time.Hour,
	})
}

func (uc userUseCase) GetUser(ctx context.Context, email string) (*domain.User, error) {
	user, err := uc.userDAO.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, custom_errors.NewNotFoundError("user not found")
	}

	return user, nil
}

func (uc userUseCase) GetUsersList(ctx context.Context) ([]*domain.User, error) {
	return uc.userDAO.GetUsers(ctx)
}

func (uc userUseCase) UpdateUserActiveStatus(ctx context.Context, params UpdateUserActiveParams) error {
	user, err := uc.userDAO.GetById(ctx, params.UserId)
	if err != nil {
		return err
	}
	if user.IsActive == params.IsActive {
		return custom_errors.NewUpdateError("The user already has this status")
	}
	user.IsActive = params.IsActive
	return uc.userDAO.Update(ctx, user)
}

func (uc userUseCase) UpdateUserRole(ctx context.Context, params UpdateUserRoleParams) error {
	user, err := uc.userDAO.GetById(ctx, params.UserId)
	if err != nil {
		return err
	}
	if user.Role.Name == params.Role {
		return custom_errors.NewUpdateError("The user already has this role")
	}
	switch params.Role {
	case domain.USER, domain.MANAGER, domain.ADMIN:
		role, err := uc.roleDAO.GetByName(ctx, params.Role)
		if err != nil {
			return err
		}
		user.RoleId = role.Id
	default:
		return custom_errors.NewUpdateError("Wrong role name")
	}
	return uc.userDAO.Update(ctx, user)
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

type UpdateUserActiveParams struct {
	UserId   uuid.UUID
	IsActive bool
}

type UpdateUserRoleParams struct {
	UserId uuid.UUID
	Role   string
}
