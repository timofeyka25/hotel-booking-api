package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"hotel-booking-app/internal/handler/dto"
	"hotel-booking-app/internal/usecase"
	"hotel-booking-app/pkg/custom_errors"
)

type PaymentHandler struct {
	validator      *validator.Validate
	paymentUseCase usecase.PaymentUseCase
}

func NewPaymentHandler(
	paymentUseCase usecase.PaymentUseCase,
	validator *validator.Validate,
) *PaymentHandler {
	return &PaymentHandler{
		paymentUseCase: paymentUseCase,
		validator:      validator,
	}
}

// PayForReservation
//
// @Summary Pay for reservation
// @Description Pay for a reservation with the specified payment details.
// @Tags Payment
// @Accept json
// @Produce json
// @Param id path string true "Reservation ID"
// @Param input body dto.CreatePaymentDTO true "Payment details"
// @Security ApiKeyAuth
// @Success 200 {object} dto.ReturnIdDTO
// @Failure 400 {object} dto.ErrorDTO
// @Failure 401 {object} dto.ErrorDTO
// @Failure 500 {object} dto.ErrorDTO
// @Router /reservation/{id}/pay [post]
func (h *PaymentHandler) PayForReservation(ctx *fiber.Ctx) error {
	idDto := new(dto.GetByIdDTO)
	paymentDto := new(dto.CreatePaymentDTO)
	if err := ctx.ParamsParser(idDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := ctx.BodyParser(paymentDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(idDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(idDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	id := ctx.Cookies("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	newId, err := h.paymentUseCase.PayForReservation(ctx.Context(), toCreatePaymentParams(
		uuid.MustParse(idDto.Id),
		userId,
		paymentDto.Amount,
	))
	if err != nil {
		_, ok := err.(*custom_errors.StatusError)
		if ok {
			return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(dto.ReturnIdDTO{Id: newId})
}

// GetUserPayments
//
// @Summary Get user payments
// @Description Returns a list of payments made by the authenticated user.
// @Tags Payment
// @Accept json
// @Produce json
// @Success 200 {array} dto.PaymentDTO
// @Failure 400 {object} dto.ErrorDTO
// @Failure 401 {object} dto.ErrorDTO
// @Failure 500 {object} dto.ErrorDTO
// @Router /payment/all [get]
func (h *PaymentHandler) GetUserPayments(ctx *fiber.Ctx) error {
	id := ctx.Cookies("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	payments, err := h.paymentUseCase.GetUserPayments(ctx.Context(), userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(mapDtoPayments(payments))
}
