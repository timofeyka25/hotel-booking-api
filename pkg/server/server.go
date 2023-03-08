package server

import "github.com/gofiber/fiber/v2"

func NewHTTPServer(cfg Config) *fiber.App {
	app := fiber.New(FiberConfig(cfg))

	return app
}
