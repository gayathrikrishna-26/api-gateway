package main

import (
	"net/http"

	"github.com/gayathrikrishna/api-gateway/handlers"
	"github.com/gayathrikrishna/api-gateway/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	r.Run(":8080")
}
