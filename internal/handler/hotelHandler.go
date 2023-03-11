package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"hotel-booking-app/internal/handler/dto"
	"hotel-booking-app/internal/usecase"
)

type HotelHandler struct {
	validator    *validator.Validate
	hotelUseCase usecase.HotelUseCase
}

func NewHotelHandler(hotelUseCase usecase.HotelUseCase, validator *validator.Validate) *HotelHandler {
	return &HotelHandler{hotelUseCase: hotelUseCase, validator: validator}
}

func (h *HotelHandler) AddHotel(ctx *fiber.Ctx) error {
	addHotelDTO := new(dto.AddHotelReqDTO)
	if err := ctx.BodyParser(addHotelDTO); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.AddHotelResDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(addHotelDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.AddHotelResDTO{Message: err.Error()})
	}
	id, err := h.hotelUseCase.AddHotel(ctx.Context(), toAddHotelParams(addHotelDTO))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.AddHotelResDTO{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(dto.AddHotelResDTO{Id: id})
}
