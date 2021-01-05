package item_service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	m "github.com/uberballo/webstore/model"
)

var baseURL = "https://bad-api-assignment.reaktor.com/v2/"

func GetProducts(category string) {
	url := fmt.Sprintf("%sproducts/%s", baseURL, category)
	resp, err := http.Get(url)

	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	var products []m.Product
	if resp.StatusCode == http.StatusOK {
		err := json.NewDecoder(resp.Body).Decode(&products)
		if err != nil {
			fmt.Println(err)
		}
	}
	var m map[string]int

	m = make(map[string]int)

	for _, p := range products {
		//fmt.Print(" https://bad-api-assignment.reaktor.com/v2/availability", p.Manufacturer)
		m[p.Manufacturer] = 1
	}
	for v := range m {
		fmt.Print(" https://bad-api-assignment.reaktor.com/v2/availability/", v)
	}
}

func GetAvailability(manufacturer string) {
	url := fmt.Sprintf("%savailability/%s", baseURL, manufacturer)
	resp, err := http.Get(url)

	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Body.Read)
	var availability AvailabilityResponse
	if resp.StatusCode == http.StatusOK {
		err := json.NewDecoder(resp.Body).Decode(&availability)
		if err != nil {
			fmt.Println(err)
		}
	}
	for _, p := range availability.Response {
		fmt.Println(p.Id)
	}
}
