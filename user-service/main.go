package main

import (
	"user-service/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register user routes
	controller.RegisterUserRoutes(r)

	r.Run(":4002") // service runs on port 4002
}
