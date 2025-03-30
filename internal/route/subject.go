package route

import (
	"github.com/gin-gonic/gin"
	"github.com/killerrekt/quantstrategix/internal/controller"
	"github.com/killerrekt/quantstrategix/internal/middleware"
)

func SubjectRoute(router *gin.Engine) {
	router.POST("/subjects", middleware.Protected(), middleware.CheckRole("admin"), controller.AddSubject)
	router.GET("/subjects", controller.GetAllSubject)
}
