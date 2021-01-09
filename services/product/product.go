package product

import (
	"sort"

	"github.com/uberballo/webstore/helpers"
	m "github.com/uberballo/webstore/model"
	"github.com/uberballo/webstore/services/badapi"
)

var products = make(map[string][]m.Product)
var categories = []string{"gloves", "beanies", "facemasks"}

func checkProducts() {
	if len(products) == 0 {
		InitializeProductData()
	}
}

func getCategorysProducts(category string) []m.Product {
	return products[category]
}

func InitializeProductData() {
	productResponses, availabilityResponses := badapi.GetProductsAndAvailability(categories)
	ar := m.CombinAvailabilityResponses(availabilityResponses)
	for _, response := range productResponses {
		combined := helpers.UpdateProductsWithAvailability(response, ar)
		products[response.Category] = combined
	}
}

func init() {
	//InitializeProductData()

}

func GetProductsWithStock(category string) []m.Product {
	products := getCategorysProducts(category)
	sort.Sort(m.ByName(products))
	return products
}
