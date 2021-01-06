package product

import (
	"fmt"

	"github.com/uberballo/webstore/helpers"
	m "github.com/uberballo/webstore/model"
	"github.com/uberballo/webstore/services/badapi"
)

var products = make(map[string][]m.Product)
var Availability []m.Availability
var categories = []string{"gloves", "beanies", "facemasks"}

func checkProducts() {
	if len(products) == 0 {
		initializeProductData()
	}
}

func getCategorysProducts(category string) []m.Product {
	checkProducts()
	return products[category]
}

func initializeProductData() {
	productResponses, availabilityResponses := badapi.GetProductsAndAvailability(categories)
	ar := m.CombinAvailabilityResponses(availabilityResponses)
	for _, response := range productResponses {
		combined := helpers.UpdateProductsWithAvailability(response, ar)
		products[response.Category] = combined
	}
}

func init() {
	fmt.Println("Updating")
	initializeProductData()
	fmt.Println("done!")
}

func GetProductsWithStock(category string) []m.Product {
	return getCategorysProducts(category)
}
