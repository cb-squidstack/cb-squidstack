
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/inventory/:productId", func(c *gin.Context) {
        productId := c.Param("productId")
        c.JSON(200, gin.H{
            "productId": productId,
            "status":    "available",
        })
    })

    r.POST("/inventory/update", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "inventory updated"})
    })

    r.GET("/health", func(c *gin.Context) {
        c.String(200, "OK")
    })

    r.Run(":8080")
}
