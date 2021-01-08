package v1

import (
	"fmt"
	"net/http"

	service "github.com/uberballo/webstore/services/product"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	category := c.Param("category")
	fmt.Println(category)
	data := service.GetProductsWithStock(category)
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   data,
	})
}
