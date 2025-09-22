package main

import (
	"product-service/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register product routes
	controller.RegisterProductRoutes(r)

	r.Run(":4003") // product-service runs on port 4003
}
