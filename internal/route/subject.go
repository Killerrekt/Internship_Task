package route

import (
	"github.com/gin-gonic/gin"
	"github.com/killerrekt/quantstrategix/internal/controller"
)

func SubjectRoute(router *gin.Engine) {
	router.POST("/subjects", controller.AddSubject)
	router.GET("/subjects", controller.GetAllSubject)
}
