package product

import (
	"sort"

	"github.com/uberballo/warehouse/server/helpers"
	m "github.com/uberballo/warehouse/server/model"
	"github.com/uberballo/warehouse/server/services/badapi"
)

var products = make(map[string][]m.Product)
var categories = []string{"gloves", "beanies", "facemasks"}

func productsExist() bool {
	return len(products) == 0
}

func getCategorysProducts(category string) []m.Product {
	return products[category]
}

//InitializeProductData initializes the products with data from Bad api
func InitializeProductData() {
	badAPIResponse := badapi.GetProductsAndAvailability(categories)

	availabilityResponse := m.CombineAvailabilityResponses(badAPIResponse.AvailabilityResponses)
	for _, productResponse := range badAPIResponse.ProductResponses {
		combined := helpers.UpdateProductsWithAvailability(productResponse, availabilityResponse)
		products[productResponse.Category] = combined
	}
}

//GetProductsWithStock returns products from the given category
func GetProductsWithStock(category string) []m.Product {
	products := getCategorysProducts(category)
	sort.Sort(m.ByName(products))
	return products
}
