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
	r.Run(":3000") // Gateway runs on port 3000
}
