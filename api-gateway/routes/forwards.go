package routes

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func forwardToAuthService(c *gin.Context) {
	resp, err := http.Post("http://localhost:3001/auth/login", "application/json", c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Auth service dfsfffsgf"})
		//dkhfk;aljdflkxcc
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func forwardToUserService(c *gin.Context) {
	req, _ := http.NewRequest("GET", "http://localhost:3002/user/profile", nil)
	req.Header = c.Request.Header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "User service unavailable"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}
