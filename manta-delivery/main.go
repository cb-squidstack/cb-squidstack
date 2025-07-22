package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/ship", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "shipment initiated"})
	})

	r.GET("/delivery/:orderId", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"delivery": "delivery status"})
	})

	r.Run(":8080")
}
