package jwt

import (
	"github.com/golang-jwt/jwt"
	"hotel-booking-app/pkg/custom_errors"
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
		return nil, custom_errors.NewUnauthorizedError()
	}

	if !parsed.Valid {
		return nil, custom_errors.NewUnauthorizedError()
	}

	return &TokenParsedParams{Id: claims.Id, Role: claims.Role}, nil
}
