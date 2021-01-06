package itemservice

import (
	"github.com/uberballo/webstore/helpers"
	. "github.com/uberballo/webstore/model"
)

func GetProductsWithStock(categories []string) []Product {
	productsResponse, availabilityResponse := GetProductsAndAvailability(categories)
	pr := CombineProductResponses(productsResponse)
	ar := CombinAvailabilityResponses(availabilityResponse)
	products := helpers.UpdateProductsWithAvailability(pr, ar)
	return products
}
