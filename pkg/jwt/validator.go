package jwt

import (
	"github.com/golang-jwt/jwt"
	"hotel-booking-app/pkg/customErrors"
)

type TokenValidator struct {
	secretKey string
}

func NewTokenValidator(cfg Config) *TokenValidator {
	return &TokenValidator{
		secretKey: cfg.SecretKey,
	}
}

func (v *TokenValidator) ValidateToken(token string) error {
	parsed, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		return []byte(v.secretKey), nil
	})
	if err != nil {
		return customErrors.NewUnauthorizedError()
	}

	if !parsed.Valid {
		return customErrors.NewUnauthorizedError()
	}

	return nil
}
