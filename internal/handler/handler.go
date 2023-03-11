package handler

import "github.com/gofiber/fiber/v2"

type handler struct {
	userHandler  *UserHandler
	hotelHandler *HotelHandler
}

func NewHandler(userHandler *UserHandler, hotelHandler *HotelHandler) *handler {
	return &handler{userHandler: userHandler, hotelHandler: hotelHandler}
}

func (h *handler) InitRoutes(app *fiber.App) {
	app.Get("/sign-in", h.userHandler.SignIn)
	app.Post("/sign-up", h.userHandler.SignUp)
	app.Post("/hotel", h.hotelHandler.AddHotel)
}
