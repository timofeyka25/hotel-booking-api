package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"hotel-booking-app/internal/handler/dto"
	"hotel-booking-app/internal/usecase"
	"hotel-booking-app/pkg/customErrors"
)

type PaymentHandler struct {
	validator      *validator.Validate
	paymentUseCase usecase.PaymentUseCase
}

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
		_, ok := err.(*customErrors.StatusError)
		if ok {
			return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(dto.ReturnIdDTO{Id: newId})
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
