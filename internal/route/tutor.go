package route

import (
	"github.com/gin-gonic/gin"
	"github.com/killerrekt/quantstrategix/internal/controller"
	"github.com/killerrekt/quantstrategix/internal/middleware"
)

func TutorRoute(router *gin.Engine) {
	router.POST("/tutor/profile", middleware.Protected(), middleware.CheckRole("tutor"), controller.TutorDetails)
	router.GET("/tutor", controller.GetTutorBySubjects)
}
