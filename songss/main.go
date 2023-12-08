package main

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
)

const (
	port      = 8080
	songsDir  = "bachgavotteshort"
	videoDir  = "footballmp4"
	mediaHTML = "templates/player.html"
)

// MediaData represents data to be passed to the HTML template.
type MediaData struct {
	VideoURL string
	AudioURL string
}

func main() {
	r := gin.Default()

	// Middleware for CORS support
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	// Route for serving media player page
	r.GET("/play/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		videoURL := fmt.Sprintf("/video/%s", filename)
		audioURL := fmt.Sprintf("/audio/%s", filename)

		data := MediaData{
			VideoURL: videoURL,
			AudioURL: audioURL,
		}

		render(c, mediaHTML, data)
	})

	// Route for serving video and audio files
	r.Static("/video", videoDir)
	r.Static("/audio", songsDir)

	fmt.Printf("Starting server on %v\n", port)

	// Run the server
	if err := r.Run(fmt.Sprintf(":%v", port)); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func render(c *gin.Context, templatePath string, data interface{}) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}

	err = tmpl.Execute(c.Writer, data)
	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}
}

