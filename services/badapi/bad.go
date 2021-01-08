package badapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	. "github.com/uberballo/webstore/model"
)

type Response struct {
	err    error
	Result interface{}
}

var baseURL = "https://bad-api-assignment.reaktor.com/v2/"

func fetchAvailability(manufacturer string, ch chan<- Response) {
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

					ch <- Response{
						err:    fmt.Errorf("Failed to fetch %s availability", manufacturer),
						Result: nil,
					}
				}
				goto RETRY
			}
			ch <- Response{
				err:    nil,
				Result: availability}
		}
	}
}

func fetchProducts(category string, ch chan<- Response) {
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
					ch <- Response{
						err:    fmt.Errorf("Failed to fetch %s catalog ", category),
						Result: nil,
					}
				}
				goto RETRY
			}
		}
		ch <- Response{
			err: nil,
			Result: ProductResponse{
				Category: category,
				Response: products}}
	}
}

func getProductsWithoutStock(categories []string) []ProductResponse {
	productChannel := make(chan Response)
	var result []ProductResponse

	for _, item := range categories {
		go fetchProducts(item, productChannel)
	}

	for range categories {
		res := <-productChannel
		result = append(result, res.Result.(ProductResponse))
	}

	return result
}

func getAvailability(manufacturers []string) []AvailabilityResponse {
	availabilityChannel := make(chan Response)
	var result []AvailabilityResponse

	for _, manu := range manufacturers {
		fmt.Println(manu, len(manu))
		go fetchAvailability(manu, availabilityChannel)
	}

	for range manufacturers {
		res := <-availabilityChannel
		if res.err != nil {
			fmt.Println(res.err)
		}
		result = append(result, res.Result.(AvailabilityResponse))
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

func createSliceOfManufacturers(m map[string]bool) []string {
	result := []string{}
	for manufacturer := range m {
		result = append(result, manufacturer)
	}
	return result
}

func GetProductsAndAvailability(categories []string) ([]ProductResponse, []AvailabilityResponse) {
	start := time.Now()

	productResponse := getProductsWithoutStock(categories)
	manufacturers := createSliceOfManufacturers(createManufacturersSet(productResponse))
	availabilityResponse := getAvailability(manufacturers)

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	return productResponse, availabilityResponse
}
