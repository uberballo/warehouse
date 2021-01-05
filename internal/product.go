package internal

import (
	"regexp"
	"strings"

	m "github.com/uberballo/webstore/model"
)

func getInStockValue(payload string) string {
	expression := "<INSTOCKVALUE>(.+)?</INSTOCKVALUE>"
	re := regexp.MustCompile(expression)
	return re.FindStringSubmatch(payload)[1]
}

func createAvailabilityMap(Availabilities []m.Availability) map[string]string {
	result := make(map[string]string)
	for _, availability := range Availabilities {
		key := strings.ToLower(availability.Id)
		value := getInStockValue(availability.Datapayload)
		result[key] = value
	}
	return result
}

func ReadInStockValue(availabilityMap map[string]string, products []m.ProductWithoutStock) []m.Product {

	return []m.Product{}
}
