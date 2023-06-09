package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/skip"
	"hotel-booking-app/pkg/jwt"
	"regexp"
)

func NewHTTPServer(jwtValidator *jwt.TokenValidator) *fiber.App {
	app := fiber.New(FiberConfig())
	validatorMiddleware := NewTokenValidatorMiddleware(jwtValidator)
	app.Use(
		cors.New(cors.Config{
			AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin", //nolint: lll
			AllowOrigins:     "*",
			AllowCredentials: true,
			AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		}),
		logger.New(),
		skip.New(validatorMiddleware.validateToken, func(ctx *fiber.Ctx) bool {
			path := ctx.Path()
			isSwaggerPath, _ := regexp.MatchString("/swagger/([a-z]*)", path)
			isAuthPath, _ := regexp.MatchString("/sign-([a-z]*)", path)

			return isAuthPath || isSwaggerPath
		}),
	)
	return app
}
