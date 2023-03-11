package handler

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	_ "hotel-booking-app/docs"
)

type handler struct {
	userHandler  *UserHandler
	hotelHandler *HotelHandler
}

func NewHandler(userHandler *UserHandler,
	hotelHandler *HotelHandler) *handler {
	return &handler{
		userHandler:  userHandler,
		hotelHandler: hotelHandler,
	}
}

func (h *handler) InitRoutes(app *fiber.App) {
	// auth handlers
	app.Get("/sign-in", h.userHandler.SignIn)
	app.Post("/sign-up", h.userHandler.SignUp)

	// hotel handlers
	app.Post("/hotel", h.hotelHandler.AddHotel)
	app.Get("/hotel", h.hotelHandler.GetAllHotels)
	app.Get("/hotel/:id", h.hotelHandler.GetHotelById)
	app.Put("/hotel/:id", h.hotelHandler.UpdateHotel)
	app.Delete("/hotel/:id", h.hotelHandler.DeleteHotel)

	// swagger handler
	app.Get("/swagger/*", swagger.HandlerDefault)
}
