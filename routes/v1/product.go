package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uberballo/webstore/services/itemservice"
)

func GetProducts(c *gin.Context) {
	categories := []string{"beanies", "gloves", "masks"}
	products := itemservice.GetProductsWithStock(categories)
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   products,
	})
}
