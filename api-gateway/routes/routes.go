package routes

import (
	// "api-gateway/middleware"

	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Gateway is aliveeee!!!"})
	})

	// Auth forwarding
	authTarget, _ := url.Parse("http://localhost:4001") // auth-service
	authProxy := httputil.NewSingleHostReverseProxy(authTarget)

	auth := r.Group("/auth")
	{
		// forward ANY /auth/* to auth-service
		auth.Any("/*path", func(c *gin.Context) {
			authProxy.ServeHTTP(c.Writer, c.Request)
		})
	}

	// User forwarding (protected)
	userTarget, _ := url.Parse("http://user-service:4002") // user-service
	userProxy := httputil.NewSingleHostReverseProxy(userTarget)

	user := r.Group("/user")
	// user.Use(middleware.AuthMiddleware())
	{
		// forward ANY /user/* to user-service
		user.Any("/*path", func(c *gin.Context) {
			userProxy.ServeHTTP(c.Writer, c.Request)
		})
	}

	productTarget, _ := url.Parse("http://localhost:4003") //product-service
	productProxy := httputil.NewSingleHostReverseProxy((productTarget))

	product := r.Group("/product")
	// user.Use(middleware.AuthMiddleware())
	{
		product.Any("/*path", func(c *gin.Context) {
			productProxy.ServeHTTP(c.Writer, c.Request)
		})
	}

	adminTarget, _ := url.Parse("http://localhost:4004") //admin-service
	adminProxey := httputil.NewSingleHostReverseProxy(adminTarget)
	admin := r.Group("/admin")
	// admin.Use(middleware.AuthMiddleware()())
	{
		admin.Any("/*path", func(c *gin.Context) {
			adminProxey.ServeHTTP(c.Writer, c.Request)
		})
	}

}
