package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

// AddHotel
//
//	@Summary	Add a new hotel
//	@Tags		Hotel
//	@Accept		json
//	@Param		input	body dto.AddHotelDTO true "Hotel data"
//	@Success	201		{object}	dto.ReturnIdDTO
//	@Failure	400		{object}	dto.ErrorDTO
//	@Failure	500		{object}	dto.ErrorDTO
//	@Router		/hotel [post]
func (h *HotelHandler) AddHotel(ctx *fiber.Ctx) error {
	addHotelDTO := new(dto.AddHotelDTO)
	if err := ctx.BodyParser(addHotelDTO); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(addHotelDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	id, err := h.hotelUseCase.AddHotel(ctx.Context(), toAddHotelParams(addHotelDTO))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(dto.ReturnIdDTO{Id: id})
}

// GetAllHotels
//
//	@Summary	Get all hotels
//	@Tags		Hotel
//	@Accept		json
//	@Success	200		{object}	[]dto.HotelDTO
//	@Failure	500		{object}	dto.ErrorDTO
//	@Router		/hotel/all [get]
func (h *HotelHandler) GetAllHotels(ctx *fiber.Ctx) error {
	hotels, err := h.hotelUseCase.GetAllHotels(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(mapDtoHotels(hotels))
}

// GetHotelById
//
//	@Summary	Get a hotel by ID
//	@Tags		Hotel
//	@Accept		json
//	@Param		id		path	string	true	"Hotel ID"
//	@Success	200		{object}	dto.HotelDTO
//	@Failure	400		{object}	dto.ErrorDTO
//	@Failure	500		{object}	dto.ErrorDTO
//	@Router		/hotel/{id} [get]
func (h *HotelHandler) GetHotelById(ctx *fiber.Ctx) error {
	idDto := new(dto.GetByIdDTO)
	if err := ctx.ParamsParser(idDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(idDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	hotel, err := h.hotelUseCase.GetHotelById(ctx.Context(), uuid.MustParse(idDto.Id))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(mapDtoHotel(hotel))
}

// UpdateHotel
//
// @Summary Update hotel by id
// @Tags Hotel
// @Accept json
// @Param id path string true "Hotel ID"
// @Param input body dto.UpdateHotelDTO true "Hotel data"
// @Success 200 {object} dto.SuccessDTO
// @Failure 400 {object} dto.ErrorDTO
// @Failure 500 {object} dto.ErrorDTO
// @Router /hotel/{id} [put]
func (h *HotelHandler) UpdateHotel(ctx *fiber.Ctx) error {
	idDto := new(dto.GetByIdDTO)
	updateDto := new(dto.UpdateHotelDTO)
	if err := ctx.ParamsParser(idDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := ctx.BodyParser(updateDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(idDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(updateDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.hotelUseCase.UpdateHotel(ctx.Context(), toUpdateHotelParams(uuid.MustParse(idDto.Id), updateDto)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(dto.SuccessDTO{Message: "Updated"})
}

// DeleteHotel
//
// @Summary Delete a hotel by ID
// @Tags Hotel
// @Param id path string true "Hotel ID"
// @Success 200 {object} dto.SuccessDTO
// @Failure 400 {object} dto.ErrorDTO
// @Failure 500 {object} dto.ErrorDTO
// @Router /hotel/{id} [delete]
func (h *HotelHandler) DeleteHotel(ctx *fiber.Ctx) error {
	idDto := new(dto.GetByIdDTO)
	if err := ctx.ParamsParser(idDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(idDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.hotelUseCase.DeleteHotel(ctx.Context(), uuid.MustParse(idDto.Id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(dto.SuccessDTO{Message: "Deleted"})
}
