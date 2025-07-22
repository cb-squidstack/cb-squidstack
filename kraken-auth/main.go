package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		// Stub logic for login
		c.JSON(http.StatusOK, gin.H{"status": "login successful"})
	})

	r.POST("/register", func(c *gin.Context) {
		// Stub logic for registration
		c.JSON(http.StatusOK, gin.H{"status": "registration successful"})
	})

	r.GET("/me", func(c *gin.Context) {
		// Stub logic to get user info
		c.JSON(http.StatusOK, gin.H{"user": "user data"})
	})

	r.Run(":8080") // Run on default port :8080
}
