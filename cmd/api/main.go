package main

import (
	"github.com/gin-gonic/gin"
	"github.com/killerrekt/quantstrategix/internal/database"
)

func main() {
	database.ConnectDB()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}
