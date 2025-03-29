package route

import (
	"github.com/gin-gonic/gin"
	"github.com/killerrekt/quantstrategix/internal/controller"
	"github.com/killerrekt/quantstrategix/internal/middleware"
)

func BookingRoute(router *gin.Engine) {
	router.POST("/bookings", middleware.Protected(), controller.BookTutor)
	router.GET("/bookings", middleware.Protected(), controller.GetBooking)
	router.PATCH("/bookings", controller.UpdateBooking)
}
