package product

import (
	"sort"

	"github.com/uberballo/warehouse/server/helpers"
	m "github.com/uberballo/warehouse/server/model"
	"github.com/uberballo/warehouse/server/services/badapi"
)

var products = make(map[string][]m.Product)
var categories = []string{"gloves", "beanies", "facemasks"}
var failedCategories = []string{}

func productsExist() bool {
	return len(products) == 0
}

func getCategorysProducts(category string) []m.Product {
	return products[category]
}

func InitializeProductData() {
	badApiResponse := badapi.GetProductsAndAvailability(categories)
	ar := m.CombinAvailabilityResponses(badApiResponse.AvailabilityResponses)
	for _, response := range badApiResponse.ProductResponses {
		combined := helpers.UpdateProductsWithAvailability(response, ar)
		products[response.Category] = combined
	}
	failedCategories = badApiResponse.AvailabilityErrors
}

func GetProductsWithStock(category string) []m.Product {
	products := getCategorysProducts(category)
	sort.Sort(m.ByName(products))
	return products
}

func GetAvailabilityErrors() []string {
	return failedCategories
}
