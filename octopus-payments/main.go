package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/pay", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "payment processed"})
	})

	r.GET("/payment/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"payment": "payment details"})
	})

	r.Run(":8080")
}
