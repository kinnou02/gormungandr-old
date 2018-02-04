package main

import (
    "github.com/gin-gonic/gin"
    "github.com/kinnou02/gormungandr/handlers"
    "github.com/kinnou02/gormungandr/journeys"
)


func setupRouter() *gin.Engine {
    // Disable Console Color
    // gin.DisableConsoleColor()
    r := gin.Default()

    r.GET("/status", handlers.Index)

    r.GET("/journeys", journeys.JourneysHandler)

    return r
}

func main() {
    r := setupRouter()
    // Listen and Server in 0.0.0.0:8080
    r.Run(":8080")
}
