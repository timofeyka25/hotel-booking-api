package handler

import "github.com/gofiber/fiber/v2"

type handler struct {
	uh *UserHandler
}

func NewHandler(uh *UserHandler) *handler {
	return &handler{uh: uh}
}

func (h *handler) InitRoutes(app *fiber.App) {
	app.Get("/sign-in", h.uh.SignIn)
}
