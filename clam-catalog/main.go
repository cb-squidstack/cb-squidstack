package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"products": "list of products"})
	})

	r.POST("/products", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{"status": "product created"})
	})

	r.GET("/products/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"product": "product details"})
	})

	r.Run(":8080")
}
