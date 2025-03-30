package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/killerrekt/quantstrategix/internal/model"
)

type CustomClaims struct {
	model.TokenPayload
	jwt.RegisteredClaims
}

func extractTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is missing")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("invalid Authorization header format")
	}

	return parts[1], nil
}

func Protected() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := extractTokenFromHeader(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.Response{Message: err.Error(), Status: false})
			c.Abort()
			return
		}
		token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.Response{Message: err.Error(), Status: false})
			c.Abort()
			return
		}
		fmt.Println(token.Claims)
		if claims, ok := token.Claims.(jwt.MapClaims)["sub"]; ok && token.Valid {
			c.Set("email", claims)
		} else {
			c.JSON(http.StatusUnauthorized, model.Response{Message: "invalid token", Status: false})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims)["role"]; ok && token.Valid {
			c.Set("role", claims)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, model.Response{Message: "invalid token", Status: false})
			c.Abort()
			return
		}
	}
}
