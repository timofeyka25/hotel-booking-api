package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"hotel-booking-app/internal/handler/dto"
	"hotel-booking-app/internal/usecase"
	"hotel-booking-app/pkg/custom_errors"
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
//	@Failure	403		{object}	dto.ErrorDTO
//	@Failure	500		{object}	dto.ErrorDTO
//	@Router		/sign-in [post]
func (h *UserHandler) SignIn(ctx *fiber.Ctx) error {
	signInDto := new(dto.SignInRequestDTO)
	if err := ctx.BodyParser(signInDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(signInDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	token, err := h.userUseCase.SignIn(ctx.Context(), toSignInParams(signInDto))
	if err != nil {
		_, ok := err.(*custom_errors.NotActiveError)
		if ok {
			return ctx.Status(fiber.StatusForbidden).JSON(dto.ErrorDTO{Message: err.Error()})
		}
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
func (h *UserHandler) SignUp(ctx *fiber.Ctx) error {
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

// GetUsersList
//
// @Summary Get users list (for admin)
// @Description Returns a list of all users.
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {array} dto.UserDTO
// @Failure 401 {object} dto.ErrorDTO
// @Failure 500 {object} dto.ErrorDTO
// @Router /users [get]
func (h *UserHandler) GetUsersList(ctx *fiber.Ctx) error {
	users, err := h.userUseCase.GetUsersList(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(mapDtoUsers(users))
}

// ChangeUserActive
//
// @Summary Change user active status
// @Tags User
// @Accept json
// @Param id path string true "User ID"
// @Param input body dto.IsActiveDTO true "User active status data"
// @Success 200 {object} dto.SuccessDTO
// @Failure 400 {object} dto.ErrorDTO
// @Failure 500 {object} dto.ErrorDTO
// @Router /user/:id/active [put]
func (h *UserHandler) ChangeUserActive(ctx *fiber.Ctx) error {
	idDto := new(dto.GetByIdDTO)
	isActiveDto := new(dto.IsActiveDTO)
	if err := ctx.ParamsParser(idDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := ctx.BodyParser(isActiveDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(idDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.userUseCase.UpdateUserActiveStatus(
		ctx.Context(),
		toUpdateUserStatusParams(uuid.MustParse(idDto.Id), isActiveDto.IsActive),
	); err != nil {
		_, ok := err.(*custom_errors.UpdateError)
		if ok {
			return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(dto.SuccessDTO{Message: "Updated"})
}

// ChangeUserRole
//
// @Summary Change user active status
// @Tags User
// @Accept json
// @Param id path string true "User ID"
// @Param input body dto.UpdateRoleDTO true "User role data"
// @Success 200 {object} dto.SuccessDTO
// @Failure 400 {object} dto.ErrorDTO
// @Failure 500 {object} dto.ErrorDTO
// @Router /user/:id/role [put]
func (h *UserHandler) ChangeUserRole(ctx *fiber.Ctx) error {
	idDto := new(dto.GetByIdDTO)
	if err := ctx.ParamsParser(idDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(idDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	roleDto := new(dto.UpdateRoleDTO)
	if err := ctx.BodyParser(roleDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(roleDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.userUseCase.UpdateUserRole(
		ctx.Context(),
		toUpdateRoleParams(uuid.MustParse(idDto.Id), roleDto),
	); err != nil {
		_, ok := err.(*custom_errors.UpdateError)
		if ok {
			return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(dto.SuccessDTO{Message: "Updated"})
}
