package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	// Initialize a Gin router
	server := gin.Default()

	// Define a route (GET request)
	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, Response{
			Message: "Hello world, from Go with Gin!",
		})
	})

	// Run the server on port 8081
	server.Run(":8081")
}
