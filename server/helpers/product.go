package helpers

import (
	"regexp"
	"strings"

	m "github.com/uberballo/warehouse/server/model"
	"github.com/uberballo/warehouse/server/test/mockservice"
)

var availabilityMap map[string]string

func init() {
	setAvailabilityMapTemp()
}

func getAvailabilityMap() map[string]string {
	return availabilityMap
}

func setAvailabilityMapTemp() {
	availabilityResponse := mockservice.MockAvailabilityResponse()
	availabilityMap = createAvailabilityMap(availabilityResponse.Response)
}

func getInStockValue(payload string) string {
	expression := "<INSTOCKVALUE>(.+)?</INSTOCKVALUE>"
	re := regexp.MustCompile(expression)
	found := re.FindStringSubmatch(payload)
	if len(found) > 1 {
		return found[1]
	}
	return payload
}

func createAvailabilityMap(Availabilities []m.Availability) map[string]string {
	result := make(map[string]string)
	for _, availability := range Availabilities {
		key := strings.ToLower(availability.ID)
		value := getInStockValue(availability.Datapayload)
		result[key] = value
	}
	return result
}

func createProductsWithAvailability(availabilityMap map[string]string, products []m.ProductWithoutStock) []m.Product {
	result := []m.Product{}
	for _, product := range products {
		stockInValue := availabilityMap[product.ID]
		productWithStock := m.Product{
			ProductWithoutStock: product,
			Stock:               stockInValue,
		}
		result = append(result, productWithStock)
	}
	return result
}

//UpdateProductsWithAvailability adds availability to product
func UpdateProductsWithAvailability(productResponse m.ProductResponse, availabilityResponse m.AvailabilityResponse) []m.Product {
	availabilityMap := createAvailabilityMap(availabilityResponse.Response)
	products := createProductsWithAvailability(availabilityMap, productResponse.Response)
	return products
}
