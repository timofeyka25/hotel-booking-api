package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"hotel-booking-app/handler/dto"
	"hotel-booking-app/usecase"
)

type UserHandler struct {
	validator *validator.Validate
	uc        usecase.UserUseCase
}

func NewUserHandler(uc usecase.UserUseCase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h UserHandler) SignIn(ctx *fiber.Ctx) error {
	return ctx.SendString("HELLO")
}

func (h UserHandler) SignUp(ctx *fiber.Ctx) error {
	signInDto := new(dto.SignUpRequestDTO)
	if err := ctx.BodyParser(signInDto); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if err := h.validator.Struct(signInDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return nil
}
