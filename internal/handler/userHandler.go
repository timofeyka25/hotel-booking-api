package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"hotel-booking-app/internal/handler/dto"
	"hotel-booking-app/internal/usecase"
)

type UserHandler struct {
	validator   *validator.Validate
	userUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase, validator *validator.Validate) *UserHandler {
	return &UserHandler{userUseCase: userUseCase, validator: validator}
}

// SignIn
//
//	@Summary	Sign in to account
//	@Tags		Authentication
//	@Accept		json
//	@Param		input	body dto.SignInRequestDTO true "User credentials"
//	@Success	200		{object}	dto.SignInResponseDTO
//	@Failure	400		{object}	dto.ErrorDTO
//	@Failure	500		{object}	dto.ErrorDTO
//	@Router		/sign-in [post]
func (h UserHandler) SignIn(ctx *fiber.Ctx) error {
	signInDto := new(dto.SignInRequestDTO)
	if err := ctx.BodyParser(signInDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(signInDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	token, err := h.userUseCase.SignIn(ctx.Context(), toSignInParams(signInDto))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}

	ctx.Cookie(&fiber.Cookie{Name: "token", Value: token})

	return ctx.JSON(dto.SignInResponseDTO{Token: token})
}

// SignUp
//
//	@Summary	Sign up into account
//	@Tags		Authentication
//	@Accept		json
//	@Param		input body dto.SignUpRequestDTO	true "User credentials"
//	@Success	201 {object}	dto.ReturnIdDTO
//	@Failure	400	{object}	dto.ErrorDTO
//	@Failure	500	{object}	dto.ErrorDTO
//	@Router		/sign-up [post]
func (h UserHandler) SignUp(ctx *fiber.Ctx) error {
	signUpDto := new(dto.SignUpRequestDTO)
	if err := ctx.BodyParser(signUpDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(signUpDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}

	id, err := h.userUseCase.SignUp(ctx.Context(), toSignUpParams(signUpDto))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(dto.ReturnIdDTO{Id: id})
}
