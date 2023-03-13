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

type TokenParsedParams struct {
	Id   string
	Role string
}

func (v *TokenValidator) ValidateToken(token string) (*TokenParsedParams, error) {
	claims := &Claims{}
	parsed, err := jwt.ParseWithClaims(token, claims, func(jwtToken *jwt.Token) (interface{}, error) {
		return []byte(v.secretKey), nil
	})
	if err != nil {
		return nil, customErrors.NewUnauthorizedError()
	}

	if !parsed.Valid {
		return nil, customErrors.NewUnauthorizedError()
	}

	return &TokenParsedParams{Id: claims.Id, Role: claims.Role}, nil
}
