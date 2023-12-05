package main

import (
    "fmt"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    // Initialize a new Gin router
    router := gin.Default()

    // Configure the songs directory name
    const songsDir = "songs"

    // Add CORS middleware
    config := cors.DefaultConfig()
    config.AllowAllOrigins = true
    router.Use(cors.New(config))

    // Add a route to handle static files
    router.Static("/BachGavotteShort", "./"+songsDir+"/BachGavotteShort")

    // Define the port
    const port = 8080

    // Start the server
    fmt.Printf("Starting server on %v\n", port)
    log.Printf("Serving %s on HTTP port: %v\n", songsDir, port)
    log.Fatal(router.Run(fmt.Sprintf(":%v", port)))
}
