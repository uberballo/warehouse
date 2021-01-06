package itemservice

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	. "github.com/uberballo/webstore/model"
)

var baseURL = "https://bad-api-assignment.reaktor.com/v2/"

func fetchAvailability(manufacturer string, ch chan<- AvailabilityResponse) {
	retryCount := 0
RETRY:
	for {
		url := fmt.Sprintf("%savailability/%s", baseURL, manufacturer)
		resp, err := http.Get(url)

		if err != nil {
			log.Panic(err)
		}
		defer resp.Body.Close()

		var availability AvailabilityResponse
		if resp.StatusCode == http.StatusOK {
			err := json.NewDecoder(resp.Body).Decode(&availability)
			if err != nil || len(availability.Response) == 0 {
				fmt.Println(err)
				retryCount++
				if retryCount > 3 {
					ch <- AvailabilityResponse{}
				}
				goto RETRY
			}
			ch <- availability
		}
	}
}

func fetchProducts(category string, ch chan<- ProductResponse) {
	retryCount := 0
RETRY:
	for {
		url := fmt.Sprintf("%sproducts/%s", baseURL, category)
		resp, err := http.Get(url)

		if err != nil {
			log.Panic(err)
		}
		defer resp.Body.Close()

		var products []ProductWithoutStock
		if resp.StatusCode == http.StatusOK {
			err := json.NewDecoder(resp.Body).Decode(&products)
			if err != nil {
				retryCount++
				if retryCount > 3 {
					ch <- ProductResponse{}
				}
				goto RETRY
			}
		}
		ch <- ProductResponse{Response: products}
	}
}

func getProductsWithoutStock(categories []string) []ProductResponse {
	productChannel := make(chan ProductResponse)
	var result []ProductResponse

	for _, item := range categories {
		go fetchProducts(item, productChannel)
	}

	for range categories {
		result = append(result, <-productChannel)
	}

	return result
}

func getAvailability(manufacturers map[string]bool) []AvailabilityResponse {
	availabilityChannel := make(chan AvailabilityResponse)
	var result []AvailabilityResponse

	for manu := range manufacturers {
		fmt.Println(manu, len(manu))
		go fetchAvailability(manu, availabilityChannel)
	}

	for range manufacturers {
		res := <-availabilityChannel
		result = append(result, res)
	}

	return result
}

func createManufacturersSet(products []ProductResponse) map[string]bool {
	manufacturers := make(map[string]bool)
	for _, response := range products {
		for _, p := range response.Response {
			manufacturers[p.Manufacturer] = true
		}
	}
	return manufacturers
}

func GetProductsAndAvailability(categories []string) ([]ProductResponse, []AvailabilityResponse) {
	start := time.Now()

	productResponse := getProductsWithoutStock(categories)
	manufacturers := createManufacturersSet(productResponse)
	availabilityResponse := getAvailability(manufacturers)

	time.Sleep(30 * time.Second)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	return productResponse, availabilityResponse
}
