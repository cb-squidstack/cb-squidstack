package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/reviews", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "review submitted"})
	})

	r.GET("/reviews/:productId", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"reviews": "product reviews"})
	})

	r.Run(":8080")
}
