package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/killerrekt/quantstrategix/internal/model"
)

func CheckRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		temp := c.GetString("role")
		if temp == role {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, model.Response{Message: "invalid role", Status: false})
			c.Abort()
			return
		}

	}
}
