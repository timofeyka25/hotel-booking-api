package handler

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	_ "hotel-booking-app/docs"
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
	app.Get("/sign-in", h.userHandler.SignIn)
	app.Post("/sign-up", h.userHandler.SignUp)

	// hotel handlers
	app.Post("/hotel", h.hotelHandler.AddHotel)
	app.Get("/hotel/all", h.hotelHandler.GetAllHotels)
	app.Get("/hotel/:id", h.hotelHandler.GetHotelById)
	app.Put("/hotel/:id", h.hotelHandler.UpdateHotel)
	app.Delete("/hotel/:id", h.hotelHandler.DeleteHotel)

	// room handlers
	app.Post("/hotel/:id/room", h.roomHandler.AddRoom)
	app.Get("/hotel/:id/room/all", h.roomHandler.GetHotelRooms)
	app.Get("/room/:id", h.roomHandler.GetRoomById)

	// swagger handler
	app.Get("/swagger/*", swagger.HandlerDefault)
}
