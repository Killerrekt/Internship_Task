package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/killerrekt/quantstrategix/internal/model"
	"github.com/killerrekt/quantstrategix/internal/query"
	"github.com/killerrekt/quantstrategix/internal/util"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var req model.RegisterReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Message: err.Error(), Status: false})
		return
	}
	err = query.SaveUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error(), Status: false})
		return
	}
	c.JSON(http.StatusAccepted, model.Response{Message: "User registered successful", Status: true})
}

func Login(c *gin.Context) {
	var req model.LoginReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Message: err.Error(), Status: false})
		return
	}

	user, err := query.FindUser(req.Email)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, model.Response{Message: "Failed to find the user", Status: false})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.Response{Message: "Invalid Password", Status: false})
		return
	}

	payload := model.TokenPayload{Email: user.Email, Role: user.Role}

	fmt.Println("pls work??", os.Getenv("JWT_SECRET"))

	access_token, err := util.CreateToken(1*time.Hour, payload, os.Getenv("JWT_SECRET"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Message: "Failed to create the token", Status: false})
		return
	}

	c.JSON(http.StatusAccepted, model.Response{Message: "Successfully logged in", Status: true, Data: model.LoginRes{AccessToken: access_token}})
}
