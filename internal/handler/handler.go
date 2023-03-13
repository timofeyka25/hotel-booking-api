package handler

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	_ "hotel-booking-app/docs"
	"hotel-booking-app/internal/domain"
	"hotel-booking-app/internal/handler/dto"
)

type handler struct {
	userHandler  *UserHandler
	hotelHandler *HotelHandler
	roomHandler  *RoomHandler
}

func NewHandler(userHandler *UserHandler,
	hotelHandler *HotelHandler,
	roomHandler *RoomHandler) *handler {
	return &handler{
		userHandler:  userHandler,
		hotelHandler: hotelHandler,
		roomHandler:  roomHandler,
	}
}

func (h *handler) InitRoutes(app *fiber.App) {
	// auth handlers
	app.Post("/sign-in", h.userHandler.SignIn)
	app.Post("/sign-up", h.userHandler.SignUp)

	// hotel handlers
	app.Post("/hotel", h.isManager, h.hotelHandler.AddHotel)
	app.Get("/hotel/all", h.hotelHandler.GetAllHotels)
	app.Get("/hotel/:id", h.hotelHandler.GetHotelById)
	app.Put("/hotel/:id", h.isManager, h.hotelHandler.UpdateHotel)
	app.Delete("/hotel/:id", h.isManager, h.hotelHandler.DeleteHotel)

	// room handlers
	app.Post("/hotel/:id/room", h.isManager, h.roomHandler.AddRoom)
	app.Get("/hotel/:id/room/all", h.roomHandler.GetHotelRooms)
	app.Get("/hotel/:id/room/free", h.roomHandler.GetHotelFreeRooms)
	app.Get("/room/:id", h.roomHandler.GetRoomById)

	// swagger handler
	app.Get("/swagger/*", swagger.HandlerDefault)
}

func (h *handler) isManager(ctx *fiber.Ctx) error {
	role := ctx.Cookies("role")
	if role == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.ErrorDTO{Message: "Unauthorized"})
	}
	if role != domain.MANAGER && role != domain.ADMIN {
		return ctx.Status(fiber.StatusForbidden).JSON(dto.ErrorDTO{Message: "Access denied"})
	}
	return ctx.Next()
}

func (h *handler) isAdmin(ctx *fiber.Ctx) error {
	role := ctx.Cookies("role")
	if role == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.ErrorDTO{Message: "Unauthorized"})
	}
	if role != domain.ADMIN {
		return ctx.Status(fiber.StatusForbidden).JSON(dto.ErrorDTO{Message: "Access denied"})
	}
	return ctx.Next()
}
