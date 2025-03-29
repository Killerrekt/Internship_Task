package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/killerrekt/quantstrategix/internal/model"
	"github.com/killerrekt/quantstrategix/internal/query"
)

func BookTutor(c *gin.Context) {
	var req model.Bookingreq

	email := c.GetString("email")

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Message: err.Error(), Status: false})
		return
	}

	data, err := query.FindUser(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error(), Status: false})
		return
	}

	entry := model.Booking{
		StudentID:   data.ID,
		TutorID:     req.TutorId,
		SubjectID:   req.Subject,
		BookingTime: req.BookingTime,
		Status:      "pending",
	}

	err = query.CreateBooking(entry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error(), Status: false})
		return
	}

	c.JSON(http.StatusAccepted, model.Response{Message: "Booking done", Status: true})
}

func GetBooking(c *gin.Context) {
	email := c.GetString("email")

	data, err := query.GetBooking(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error(), Status: false})
		return
	}

	c.JSON(http.StatusAccepted, model.Response{Message: "Booking fetched", Status: true, Data: data})
}

func UpdateBooking(c *gin.Context) {
	var req model.UpdateBookingReq

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Message: err.Error(), Status: false})
		return
	}

	err = query.UpdateBooking(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error(), Status: false})
		return
	}

	c.JSON(http.StatusAccepted, model.Response{Message: "Booking updated", Status: true})
}
