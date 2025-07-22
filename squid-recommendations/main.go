package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/recommendations/:userId", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"recommendations": "personalized suggestions"})
	})

	r.Run(":8080")
}
