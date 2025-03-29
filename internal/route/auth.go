package route

import (
	"github.com/gin-gonic/gin"
	"github.com/killerrekt/quantstrategix/internal/controller"
)

func AuthRoute(router *gin.Engine) {
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)
}
