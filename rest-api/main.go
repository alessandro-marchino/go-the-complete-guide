package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)

	// Listen on localhost:8080
	server.Run(":8080")
}

func getEvents(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ "message": "Hello!" })
}
