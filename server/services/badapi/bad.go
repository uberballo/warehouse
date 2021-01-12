package badapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/uberballo/warehouse/server/helpers/apihelper"
	. "github.com/uberballo/warehouse/server/model"
)

type response struct {
	err    error
	Result interface{}
}

type fn func(string, int, chan<- response)

var baseURL = "https://bad-api-assignment.reaktor.com/v2/"

var c = apihelper.GetHTTPClient()

func retry(f fn, param string, retryCount int, ch chan<- response, err error) {
	retryCount++
	if retryCount > 2 {
		ch <- response{
			err:    err,
			Result: nil,
		}
		return
	}
	f(param, retryCount, ch)
}

func makeRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err := errors.New("Error occurred creating new request")
		fmt.Println((err))
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func decodeAvailabilityResponse(resp http.Response) (*AvailabilityResponse, error) {
	var availability AvailabilityResponse
	if resp.StatusCode == http.StatusOK {
		err := json.NewDecoder(resp.Body).Decode(&availability)
		if err != nil {
			return nil, err
		}
		return &availability, nil
	}
	return nil, fmt.Errorf("Received status code %d", resp.StatusCode)
}

func fetchAvailabilityData(manufacturer string, retryCount int, ch chan<- response) {
	url := apihelper.CreateURL(baseURL, "availability", manufacturer)
	resp, err := makeRequest(url)
	if err != nil {
		ch <- response{
			err:    err,
			Result: nil,
		}
	}
	defer resp.Body.Close()

	availability, err := decodeAvailabilityResponse(*resp)
	if err != nil {
		err := fmt.Errorf("Failed to fetch %s's availability after %d tries", manufacturer, retryCount)
		retry(fetchAvailabilityData, manufacturer, retryCount, ch, err)
		return
	}

	ch <- response{
		err:    nil,
		Result: *availability,
	}
}

func decodeProductResponse(resp http.Response) ([]ProductWithoutStock, error) {
	var products []ProductWithoutStock
	if resp.StatusCode == http.StatusOK {
		err := json.NewDecoder(resp.Body).Decode(&products)
		if err != nil {
			return nil, err
		}
		return products, nil
	}
	return nil, fmt.Errorf("Received status code %d", resp.StatusCode)
}

func fetchProductsData(category string, retryCount int, ch chan<- response) {
	url := apihelper.CreateURL(baseURL, "products", category)
	resp, err := makeRequest(url)
	if err != nil {
		ch <- response{
			err:    err,
			Result: nil,
		}
	}
	defer resp.Body.Close()

	products, err := decodeProductResponse(*resp)
	if err != nil {
		err := fmt.Errorf("Failed to fetch %s's after %d tries", category, retryCount)
		retry(fetchProductsData, category, retryCount, ch, err)
		return
	}

	ch <- response{
		err: nil,
		Result: ProductResponse{
			Category: category,
			Response: products},
	}
}

func getProductsWithoutStock(categories []string) []ProductResponse {
	productChannel := make(chan response)
	var result []ProductResponse

	for _, item := range categories {
		go fetchProductsData(item, 0, productChannel)
	}

	for range categories {
		res := <-productChannel
		if res.err != nil {
			fmt.Println(res.err)
		}
		result = append(result, res.Result.(ProductResponse))
	}

	return result
}

func getAvailability(manufacturers []string) []AvailabilityResponse {
	availabilityChannel := make(chan response)
	var result []AvailabilityResponse

	for _, manufacturer := range manufacturers {
		go fetchAvailabilityData(manufacturer, 0, availabilityChannel)
	}

	for range manufacturers {
		res := <-availabilityChannel
		if res.err != nil {
			fmt.Println(res.err)
		} else {
			result = append(result, res.Result.(AvailabilityResponse))
		}
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

//GetProducts returns products without availability
func GetProducts(categories []string) []ProductResponse {
	productResponse := getProductsWithoutStock(categories)
	return productResponse
}

//GetProductsAndAvailability fetches each categories products and availabilities from the Bad api.
func GetProductsAndAvailability(categories []string) BadAPIResponse {
	start := time.Now()

	productResponse := getProductsWithoutStock(categories)
	manufacturers := createSliceOfManufacturers(createManufacturersSet(productResponse))
	availabilityResponse := getAvailability(manufacturers)

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	resp := BadAPIResponse{
		ProductResponses:      productResponse,
		AvailabilityResponses: availabilityResponse,
	}
	return resp
}
