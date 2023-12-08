package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	const port = 8080
	const songsDir = "footballmp4"

	// Middleware for CORS support
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	// Serve static files from the specified directory
	r.Static("/", songsDir)

	fmt.Printf("Starting server on %v\n", port)

	// Run the server
	if err := r.Run(fmt.Sprintf(":%v", port)); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
