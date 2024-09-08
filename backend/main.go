package main

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func main() {
    r := gin.Default()

    r.Use(CorsMiddleware())

    r.POST("/submit", func(c *gin.Context) {
        var json struct {
            WalletAddress string `json:"walletAddress" binding:"required"`
        }

        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Here you would typically interact with the smart contract using the walletAddress and message
        // For example, you can use go-ethereum or other blockchain libraries to perform the action.

        fmt.Printf("Received wallet address: %s\n", json.WalletAddress)

        // Assuming the smart contract execution is successful
        c.JSON(http.StatusOK, gin.H{"status": "Smart contract executed successfully"})
    })

    r.Run(":8080") // Run on port 8080
}

