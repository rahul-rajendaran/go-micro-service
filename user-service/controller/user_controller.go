package controller

import (
	"net/http"
	"user-service/service"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/list", getAllUsers)
		userGroup.POST("/create", createUser)
		userGroup.GET("/:id", getUserByID)
	}
}

func getAllUsers(c *gin.Context) {
	users := service.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func createUser(c *gin.Context) {
	var newUser service.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created := service.CreateUser(newUser)
	c.JSON(http.StatusOK, created)
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
