package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	const port = 8080
	const songsDir = "bachgavotteshort"
	const videoDir = "footballmp4"

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	r.GET("/video/:playlist.m3u8", func(c *gin.Context) {
		filename := c.Param("playlist.m3u8")
		c.File(fmt.Sprintf("%s/%s", videoDir, filename))
	})

	r.GET("/audio/:outputlist.m3u8", func(c *gin.Context) {
		filename := c.Param("outputlist.m3u8")
		c.File(fmt.Sprintf("%s/%s", songsDir, filename))
	})

	fmt.Printf("Starting server on %v\n", port)

	if err := r.Run(fmt.Sprintf(":%v", port)); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
