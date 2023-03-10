package server

import (
	"github.com/gofiber/fiber/v2"
	"hotel-booking-app/pkg/customErrors"
	"hotel-booking-app/pkg/jwt"
)

type tokenValidatorMiddleware struct {
	tokenValidator *jwt.TokenValidator
}

func NewTokenValidatorMiddleware(tokenValidator *jwt.TokenValidator) *tokenValidatorMiddleware {
	return &tokenValidatorMiddleware{tokenValidator: tokenValidator}
}

func (m *tokenValidatorMiddleware) validateToken(ctx *fiber.Ctx) error {
	if err := m.tokenValidator.ValidateToken(ctx.Cookies("token")); err != nil {
		if _, ok := err.(*customErrors.UnauthorizedError); ok {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}

		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Next()
}
