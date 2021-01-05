package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Availability struct {
	Id          string
	Datapayload string
}

type AvailabilityResponse struct {
	Response []Availability
}

var baseURL = "https://bad-api-assignment.reaktor.com/v2/"

func GetAvailability(manufacturer string, ch chan<- AvailabilityResponse) {
	retryCount := 0
RETRY:
	for {
		url := fmt.Sprintf("%savailability/%s", baseURL, manufacturer)
		resp, err := http.Get(url)
		fmt.Println(url)
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

type ProductResponse struct {
	response []Product
}

type Product struct {
	Id           string
	Name         string
	Color        []string
	Price        int
	Manufacturer string
}

func GetProducts(category string, ch chan<- ProductResponse) {
	retryCount := 0
RETRY:
	for {
		url := fmt.Sprintf("%sproducts/%s", baseURL, category)
		resp, err := http.Get(url)

		if err != nil {
			log.Panic(err)
		}
		defer resp.Body.Close()

		var products []Product
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
		ch <- ProductResponse{response: products}
	}
}

func GetStuff() ([]ProductResponse, []AvailabilityResponse) {

	start := time.Now()

	ch2 := make(chan ProductResponse)
	for _, item := range os.Args[1:] {
		go GetProducts(item, ch2)
	}
	var resp []ProductResponse
	for range os.Args[1:] {
		resp = append(resp, <-ch2)
	}

	var manufacturers map[string]int
	manufacturers = make(map[string]int)
	for _, response := range resp {
		for _, p := range response.response {
			manufacturers[p.Manufacturer] = 1
		}
	}

	ch3 := make(chan AvailabilityResponse)
	var resp2 []AvailabilityResponse
	for manu := range manufacturers {
		fmt.Println(manu, len(manu))
		go GetAvailability(manu, ch3)
	}
	for range manufacturers {
		res := <-ch3
		resp2 = append(resp2, res)
	}

	fmt.Println(len(resp))
	fmt.Println(len(manufacturers))
	fmt.Println(len(resp2))
	time.Sleep(30 * time.Second)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	return resp, resp2
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	pro, _ := GetStuff()
	f, err := os.Create("testi.txt")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, lin := range pro {
		for _, line := range lin.response {
			w.WriteString(line.Id + " " + line.Manufacturer + " " + line.Name)
		}
	}
	w.Flush()

	/*
		router := gin.Default()
		router.Use(static.Serve("/", static.LocalFile("./client/build", true)))
		api := router.Group("/api")
		{
			api.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "pong",
				})
			})
		}

		// Start and run the server
		router.Run(":5000")
	*/
}
