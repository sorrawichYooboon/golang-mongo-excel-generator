package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Define a route for the ping API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong!",
		})
	})

	// Start the server
	router.Run(":3000")
}
