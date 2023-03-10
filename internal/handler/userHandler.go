package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"hotel-booking-app/internal/handler/dto"
	"hotel-booking-app/internal/usecase"
)

type UserHandler struct {
	validator *validator.Validate
	uc        usecase.UserUseCase
}

func NewUserHandler(uc usecase.UserUseCase, v *validator.Validate) *UserHandler {
	return &UserHandler{uc: uc, validator: v}
}

func (h UserHandler) SignIn(ctx *fiber.Ctx) error {
	signInDto := new(dto.SignInRequestDTO)
	if err := ctx.BodyParser(signInDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.SignInResponseDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(signInDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.SignInResponseDTO{Message: err.Error()})
	}
	token, err := h.uc.SignIn(ctx.Context(), toSignInParams(signInDto))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.SignInResponseDTO{Message: err.Error()})
	}

	ctx.Cookies("token", token)

	return ctx.JSON(dto.SignInResponseDTO{Token: token})
}

func (h UserHandler) SignUp(ctx *fiber.Ctx) error {
	signUpDto := new(dto.SignUpRequestDTO)
	if err := ctx.BodyParser(signUpDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.SignUpResponseDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(signUpDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.SignUpResponseDTO{Message: err.Error()})
	}

	id, err := h.uc.SignUp(ctx.Context(), toSignUpParams(signUpDto))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.SignUpResponseDTO{Message: err.Error()})
	}

	return ctx.JSON(dto.SignUpResponseDTO{Id: id})
}
