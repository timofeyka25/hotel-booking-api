package handler

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	_ "hotel-booking-app/docs"
	"hotel-booking-app/internal/domain"
	"hotel-booking-app/internal/handler/dto"
)

type handler struct {
	userHandler        *UserHandler
	hotelHandler       *HotelHandler
	roomHandler        *RoomHandler
	reservationHandler *ReservationHandler
	paymentHandler     *PaymentHandler
}

func NewHandler(userHandler *UserHandler,
	hotelHandler *HotelHandler,
	roomHandler *RoomHandler,
	reservationHandler *ReservationHandler,
	paymentHandler *PaymentHandler) *handler {
	return &handler{
		userHandler:        userHandler,
		hotelHandler:       hotelHandler,
		roomHandler:        roomHandler,
		reservationHandler: reservationHandler,
		paymentHandler:     paymentHandler,
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
	app.Get("/hotel/:id/room/all", h.isManager, h.roomHandler.GetHotelRooms)
	app.Get("/hotel/:id/room/free", h.roomHandler.GetHotelFreeRooms)
	app.Get("/room/:id", h.roomHandler.GetRoomById)
	app.Put("/room/:id", h.isManager, h.roomHandler.UpdateRoom)
	app.Delete("/room/:id", h.isManager, h.roomHandler.DeleteRoom)

	// reservation handlers
	app.Post("/room/:id/reserve", h.reservationHandler.CreateReservation)
	app.Get("/reservation/all", h.reservationHandler.GetAllUserReservations)
	app.Get("/reservation/all/manager", h.isManager, h.reservationHandler.GetAllReservations)
	app.Get("/reservation/:id/cancel", h.reservationHandler.CancelUserReservation)
	app.Put("/reservation/:id/status", h.isManager, h.reservationHandler.UpdateReservationStatus)

	// payment handlers
	app.Post("/reservation/:id/pay", h.paymentHandler.PayForReservation)
	app.Get("/payment/all", h.paymentHandler.GetUserPayments)

	// admin handlers
	app.Get("/users", h.isAdmin, h.userHandler.GetUsersList)
	app.Put("/user/:id/active", h.isAdmin, h.userHandler.ChangeUserActive)
	app.Put("/user/:id/role", h.isAdmin, h.userHandler.ChangeUserRole)

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
