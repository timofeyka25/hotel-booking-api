package handler

import (
	"hotel-booking-app/internal/handler/dto"
	"hotel-booking-app/internal/usecase"
)

func toSignUpParams(dto *dto.SignUpRequestDTO) usecase.SignUpParams {
	return usecase.SignUpParams{Name: dto.Name, Email: dto.Email, Password: dto.Password}
}

func toSignInParams(dto *dto.SignInRequestDTO) usecase.SignInParams {
	return usecase.SignInParams{Email: dto.Email, Password: dto.Password}
}
