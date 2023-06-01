package server

import (
	"github.com/gofiber/fiber/v2"
	"hotel-booking-app/pkg/custom_errors"
	"hotel-booking-app/pkg/jwt"
)

type tokenValidatorMiddleware struct {
	tokenValidator *jwt.TokenValidator
}

func NewTokenValidatorMiddleware(tokenValidator *jwt.TokenValidator) *tokenValidatorMiddleware {
	return &tokenValidatorMiddleware{tokenValidator: tokenValidator}
}

func (m *tokenValidatorMiddleware) validateToken(ctx *fiber.Ctx) error {
	parsedParams, err := m.tokenValidator.ValidateToken(ctx.Cookies("token"))
	if err != nil {
		if _, ok := err.(*custom_errors.UnauthorizedError); ok {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}

		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	ctx.Cookie(&fiber.Cookie{Name: "id", Value: parsedParams.Id})
	ctx.Cookie(&fiber.Cookie{Name: "role", Value: parsedParams.Role})

	return ctx.Next()
}
