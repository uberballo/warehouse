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

func handleAvailabilityResponse(resp http.Response) (*AvailabilityResponse, error) {
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

func fetchAvailability(manufacturer string, retryCount int, ch chan<- response) {
	url := apihelper.CreateURL(baseURL, "availability", manufacturer)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err := errors.New("Error occurred creating new request")
		fmt.Println((err))
		return
	}

	//If you want to test the product if error-mode, uncomment the line.
	//req.Header.Set("x-force-error-mode", "all")
	resp, err := c.Do(req)

	if err != nil {
		err := fmt.Errorf("Error occurred during GET %s", err)
		retry(fetchAvailability, manufacturer, retryCount, ch, err)
		return
	}
	defer resp.Body.Close()

	availability, err := handleAvailabilityResponse(*resp)
	if err != nil {
		err := fmt.Errorf("Failed to fetch %s availability after %d tries", manufacturer, retryCount)
		retry(fetchAvailability, manufacturer, retryCount, ch, err)
		return
	}

	ch <- response{
		err:    nil,
		Result: *availability,
	}

}

func handleProductResponse(resp http.Response) ([]ProductWithoutStock, error) {
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

func fetchProducts(category string, retryCount int, ch chan<- response) {
	url := apihelper.CreateURL(baseURL, "products", category)
	resp, err := c.Get(url)

	if err != nil {
		err := fmt.Errorf("Error occurred during GET %s", err)
		retry(fetchProducts, category, retryCount, ch, err)
		return
	}
	defer resp.Body.Close()

	products, err := handleProductResponse(*resp)

	if err != nil {
		err := fmt.Errorf("Failed to fetch %s's after %d tries", category, retryCount)
		retry(fetchProducts, category, retryCount, ch, err)
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
		go fetchProducts(item, 0, productChannel)
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
		go fetchAvailability(manufacturer, 0, availabilityChannel)
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
