package routes

import (
	"auth-service/middleware"

	"github.com/gin-gonic/gin"
	// "net/http"
)

func RegisterRoutes(r *gin.Engine) {
	// Public route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Gateway is aliveeee!!!"})
	})

	// Auth route forwarding
	auth := r.Group("/auth")
	{
		auth.GET("/login", forwardToAuthService)
	}

	// Protected route
	user := r.Group("/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/profile", forwardToUserService)
	}
}
