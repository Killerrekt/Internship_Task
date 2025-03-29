package main

import (
	"github.com/gin-gonic/gin"
	"github.com/killerrekt/quantstrategix/internal/database"
	"github.com/killerrekt/quantstrategix/internal/route"
)

func main() {
	database.ConnectDB()
	database.RunMigration()

	router := gin.Default()

	route.AuthRoute(router)
	route.SubjectRoute(router)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}
