package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/killerrekt/quantstrategix/internal/model"
	"github.com/killerrekt/quantstrategix/internal/query"
	"gorm.io/gorm"
)

func TutorDetails(c *gin.Context) {
	email := c.GetString("email")

	var req model.AddTutorDetail
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Message: err.Error(), Status: false})
		return
	}

	check := query.CheckSubjects(req.SubjectName)
	if !check {
		c.JSON(http.StatusPreconditionFailed, model.Response{Message: "Not all subject are in DB", Status: false})
		return
	}

	var tutor model.Tutor
	tutor, err = query.GetTutor(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			data, _ := query.FindUser(email)
			tutor = model.Tutor{
				UserID:       data.ID,
				Bio:          req.Bio,
				SubjectsName: req.SubjectName,
			}

			if err := query.CreateDetail(tutor); err != nil {
				c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error(), Status: false})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error(), Status: false})
			return
		}
	} else {
		if req.Bio != "" {
			tutor.Bio = req.Bio
		}
		if len(req.SubjectName) != 0 {
			tutor.SubjectsName = append(tutor.SubjectsName, req.SubjectName...)
		}
		if err := query.AddDetail(tutor); err != nil {
			c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error(), Status: false})
			return
		}
	}

	c.JSON(http.StatusAccepted, model.Response{Message: "Details added", Status: true})
}

func GetTutorBySubjects(c *gin.Context) {
	subject := c.Query("subject")
	data, err := query.GetTutorBySubject(subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Message: "Failed to get the tutor list", Status: false})
		return
	}
	c.JSON(http.StatusAccepted, model.Response{Message: "Got the tutor list", Status: true, Data: data})
}
