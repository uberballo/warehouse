package v1

import (
	"fmt"

	"github.com/uberballo/warehouse/server/helpers/apihelper"

	service "github.com/uberballo/warehouse/server/services/product"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	category := c.Param("category")
	fmt.Println(category)
	data := service.GetProductsWithStock(category)
	body := map[string]interface{}{
		"products": data,
	}
	apihelper.Respond(c.Writer, body)
}
