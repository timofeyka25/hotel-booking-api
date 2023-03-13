package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"hotel-booking-app/internal/handler/dto"
	"hotel-booking-app/internal/usecase"
)

type RoomHandler struct {
	validator   *validator.Validate
	roomUseCase usecase.RoomUseCase
}

func NewRoomHandler(roomUseCase usecase.RoomUseCase, validator *validator.Validate) *RoomHandler {
	return &RoomHandler{roomUseCase: roomUseCase, validator: validator}
}

// AddRoom
//
// @Summary Add a new room to a hotel
// @Tags Room
// @Accept json
// @Param id path string true "Hotel ID"
// @Param input body dto.AddRoomDTO true "New room data"
// @Success 201 {object} dto.ReturnIdDTO
// @Failure 400 {object} dto.ErrorDTO
// @Failure 500 {object} dto.ErrorDTO
// @Router /hotel/{id}/room [post]
func (h *RoomHandler) AddRoom(ctx *fiber.Ctx) error {
	idDto := new(dto.GetByIdDTO)
	roomDto := new(dto.AddRoomDTO)
	if err := ctx.ParamsParser(idDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := ctx.BodyParser(roomDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(idDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(roomDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	id, err := h.roomUseCase.AddRoom(ctx.Context(), toAddRoomParams(uuid.MustParse(idDto.Id), roomDto))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(dto.ReturnIdDTO{Id: id})
}

// GetRoomById
//
// @Summary Get a room by ID
// @Tags Room
// @Accept json
// @Param id path string true "Room ID"
// @Success 200 {object} dto.RoomDTO
// @Failure 400 {object} dto.ErrorDTO
// @Failure 500 {object} dto.ErrorDTO
// @Router /room/{id} [get]
func (h *RoomHandler) GetRoomById(ctx *fiber.Ctx) error {
	idDto := new(dto.GetByIdDTO)
	if err := ctx.ParamsParser(idDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(idDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	room, err := h.roomUseCase.GetRoomById(ctx.Context(), uuid.MustParse(idDto.Id))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(mapDtoRoom(room))
}

// GetHotelRooms
//
// @Summary Get all rooms for a hotel
// @Tags Room
// @Accept json
// @Param id path string true "Hotel ID"
// @Success 200 {object} []dto.RoomDTO
// @Failure 400 {object} dto.ErrorDTO
// @Failure 500 {object} dto.ErrorDTO
// @Router /hotel/{id}/room/all [get]
func (h *RoomHandler) GetHotelRooms(ctx *fiber.Ctx) error {
	idDto := new(dto.GetByIdDTO)
	if err := ctx.ParamsParser(idDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(idDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	rooms, err := h.roomUseCase.GetHotelRooms(ctx.Context(), uuid.MustParse(idDto.Id))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(mapDtoRooms(rooms))
}

func (h *RoomHandler) GetHotelFreeRooms(ctx *fiber.Ctx) error {
	idDto := new(dto.GetByIdDTO)
	if err := ctx.ParamsParser(idDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(idDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	rooms, err := h.roomUseCase.GetHotelFreeRooms(ctx.Context(), uuid.MustParse(idDto.Id))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(mapDtoRooms(rooms))
}
