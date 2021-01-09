package v1

import (
	"fmt"

	"github.com/uberballo/webstore/helpers/apihelper"

	service "github.com/uberballo/webstore/services/product"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	category := c.Param("category")
	fmt.Println(category)
	data := service.GetProductsWithStock(category)

	apihelper.Respond(c.Writer, data)
}
