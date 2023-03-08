package usecase

import "hotel-booking-app/dao"

type UserUseCase interface {
	SignUp() error
	SignIn() error
	GetUser()
}

type userUseCase struct {
	dao dao.UserDAO
}

func NewUserUseCase(dao dao.UserDAO) *userUseCase {
	return &userUseCase{dao: dao}
}

func (uc userUseCase) SignUp() error {
	//TODO implement me
	panic("implement me")
}

func (uc userUseCase) SignIn() error {
	//TODO implement me
	panic("implement me")
}

func (uc userUseCase) GetUser() {
	//TODO implement me
	panic("implement me")
}
