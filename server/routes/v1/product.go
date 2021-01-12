package v1

import (
	"github.com/uberballo/warehouse/server/helpers/apihelper"

	service "github.com/uberballo/warehouse/server/services/product"

	"github.com/gin-gonic/gin"
)

//GetProducts fetches data and sends json with product data
func GetProducts(c *gin.Context) {
	category := c.Param("category")

	data := service.GetProductsWithStock(category)
	body := map[string]interface{}{
		"products": data,
	}
	apihelper.Respond(c.Writer, body)
}
