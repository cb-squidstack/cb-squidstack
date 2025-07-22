
package main

import (
    "io/ioutil"
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/check", func(c *gin.Context) {
        resp, err := http.Get("http://nautilus-inventory:8080/inventory/test-item")
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)
        c.String(resp.StatusCode, string(body))
    })

    r.GET("/orders", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Orders service ready"})
    })

    r.GET("/health", func(c *gin.Context) {
        c.String(200, "OK")
    })

    r.Run(":8080")
}
