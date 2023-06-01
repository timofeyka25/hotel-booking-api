package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"hotel-booking-app/internal/handler/dto"
	"hotel-booking-app/internal/usecase"
	"hotel-booking-app/pkg/custom_errors"
)

type ReservationHandler struct {
	validator          *validator.Validate
	reservationUseCase usecase.ReservationUseCase
}

func NewReservationHandler(
	reservationUseCase usecase.ReservationUseCase,
	validator *validator.Validate) *ReservationHandler {
	return &ReservationHandler{
		validator:          validator,
		reservationUseCase: reservationUseCase,
	}
}

// CreateReservation
//
// @Summary Create a reservation for a room in a hotel
// @Tags Reservation
// @Accept json
// @Param id path string true "Room ID"
// @Param input body dto.CreateReservationDTO true "Reservation data"
// @Success 201 {object} dto.ReturnIdDTO
// @Failure 400 {object} dto.ErrorDTO
// @Failure 401 {object} dto.ErrorDTO
// @Failure 500 {object} dto.ErrorDTO
// @Router /room/{id}/reserve [post]
func (h *ReservationHandler) CreateReservation(ctx *fiber.Ctx) error {
	roomIdDto := new(dto.GetByIdDTO)
	reservationDto := new(dto.CreateReservationDTO)
	if err := ctx.ParamsParser(roomIdDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := ctx.BodyParser(reservationDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(roomIdDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	parsedDto, err := reservationDto.ParseAndValidate()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	userId := ctx.Cookies("id")
	if userId == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.ErrorDTO{
			Message: custom_errors.NewUnauthorizedError().Error()})
	}
	id, err := h.reservationUseCase.CreateReservation(ctx.Context(),
		toCreateReservationParams(uuid.MustParse(roomIdDto.Id), uuid.MustParse(userId), parsedDto))
	if err != nil {
		_, ok := err.(*custom_errors.AlreadyReservedError)
		if ok {
			return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(dto.ReturnIdDTO{Id: id})
}

// GetAllUserReservations
//
// @Summary Get all reservations for the authenticated user
// @Tags Reservation
// @Accept json
// @Produce json
// @Success 200 {array} dto.ReservationDTO
// @Failure 400 {object} dto.ErrorDTO
// @Failure 500 {object} dto.ErrorDTO
// @Router /reservation/all [get]
func (h *ReservationHandler) GetAllUserReservations(ctx *fiber.Ctx) error {
	id := ctx.Cookies("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	reservations, err := h.reservationUseCase.GetAllUserReservations(ctx.Context(), userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(mapDtoReservations(reservations))
}

// GetAllReservations
//
// @Summary Get all reservations (this request for the manager)
// @Tags Reservation
// @Accept json
// @Produce json
// @Success 200 {array} dto.ReservationDTO
// @Failure 400 {object} dto.ErrorDTO
// @Failure 401 {object} dto.ErrorDTO
// @Failure 500 {object} dto.ErrorDTO
// @Router /reservation/all/manager [get]
func (h *ReservationHandler) GetAllReservations(ctx *fiber.Ctx) error {
	reservations, err := h.reservationUseCase.GetAllReservations(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(mapDtoReservations(reservations))
}

// CancelUserReservation
//
// @Summary Cancel a user's reservation
// @Tags Reservation
// @Accept json
// @Param id path string true "Reservation ID"
// @Success 200 {object} dto.SuccessDTO
// @Failure 400 {object} dto.ErrorDTO
// @Failure 401 {object} dto.ErrorDTO
// @Failure 500 {object} dto.ErrorDTO
// @Router /reservation/:id/cancel [get]
func (h *ReservationHandler) CancelUserReservation(ctx *fiber.Ctx) error {
	idDto := new(dto.GetByIdDTO)
	if err := ctx.ParamsParser(idDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(idDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	id := ctx.Cookies("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	err = h.reservationUseCase.CancelUserReservation(ctx.Context(), uuid.MustParse(idDto.Id), userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(dto.SuccessDTO{Message: "Cancelled"})
}

// UpdateReservationStatus
//
// @Summary Update the status of a reservation
// @Tags Reservation
// @Accept json
// @Param id path string true "Reservation ID"
// @Param input body dto.UpdateReservationStatusDTO true "Reservation status data"
// @Success 200 {object} dto.SuccessDTO
// @Failure 400 {object} dto.ErrorDTO
// @Failure 500 {object} dto.ErrorDTO
// @Router /reservation/:id/status [put]
func (h *ReservationHandler) UpdateReservationStatus(ctx *fiber.Ctx) error {
	idDto := new(dto.GetByIdDTO)
	updateDto := new(dto.UpdateReservationStatusDTO)
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
	if err := h.reservationUseCase.UpdateStatus(ctx.Context(), toUpdateReservationStatusParams(
		uuid.MustParse(idDto.Id),
		updateDto,
	)); err != nil {
		_, ok := err.(*custom_errors.StatusError)
		if ok {
			return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(dto.SuccessDTO{Message: "Updated"})
}
