package main

import (
	"api-gateway/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)
	fmt.Println("Server Started")
	r.Run(":8080")
}
