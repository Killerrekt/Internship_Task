package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/killerrekt/quantstrategix/internal/model"
	"github.com/killerrekt/quantstrategix/internal/query"
)

func AddSubject(c *gin.Context) {
	var req model.AddSubjectReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Message: err.Error(), Status: false})
		return
	}

	err = query.AddSubject(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Message: "Failed to add the subject", Status: false})
		return
	}

	c.JSON(http.StatusAccepted, model.Response{Message: "Subject added", Status: true})
}

func GetAllSubject(c *gin.Context) {
	data, err := query.GetAllSubject()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Message: "Failed to get all the subjects", Status: false})
		return
	}

	c.JSON(http.StatusAccepted, model.Response{Message: "All subject retrived", Status: true, Data: data})
}
