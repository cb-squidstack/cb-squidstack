package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/metrics", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"metrics": "prometheus metrics"})
	})

	r.GET("/analytics/dashboard", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"dashboard": "system analytics"})
	})

	r.Run(":8080")
}
