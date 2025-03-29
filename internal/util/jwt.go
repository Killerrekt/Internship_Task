package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/killerrekt/quantstrategix/internal/model"
)

func CreateToken(exp time.Duration, payload model.TokenPayload, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(exp).Unix(),
	}

	claims["sub"] = payload.Email
	claims["role"] = payload.Role

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, err
}
