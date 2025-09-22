package controller

import (
	"net/http"
	"product-service/service"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	productGroup := r.Group("/product")
	{
		productGroup.GET("/list", getAllProducts)
		productGroup.POST("/create", createProduct)
		productGroup.GET("/:id", getProductByID)
	}
}

func getAllProducts(c *gin.Context) {
	products := service.GetAllProducts()
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func createProduct(c *gin.Context) {
	var newProduct service.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created := service.CreateProduct(newProduct)
	c.JSON(http.StatusOK, created)
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")
	product, err := service.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}
