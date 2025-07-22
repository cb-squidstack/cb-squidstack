package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/notify", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "notification sent"})
	})

	r.GET("/notifications/:userId", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"notifications": "user notifications"})
	})

	r.Run(":8080")
}
