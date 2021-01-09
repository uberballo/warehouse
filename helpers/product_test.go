package helpers

import (
	"reflect"
	"sort"
	"testing"

	"github.com/uberballo/webstore/test/mock_service"

	. "github.com/uberballo/webstore/model"
)

func getCorrectProducts() []Product {
	var correctProducts = []Product{
		Product{
			ProductWithoutStock: ProductWithoutStock{
				Id:           "ae8c8ad79a3e4a554d6f2",
				Type:         "beanies",
				Name:         "SOPREV STAR",
				Color:        []string{"purple"},
				Price:        55,
				Manufacturer: "umpante",
			},
			Stock: "OUTOFSTOCK",
		},
		Product{
			ProductWithoutStock: ProductWithoutStock{
				Id:           "3af7caee9be9365e49e93576",
				Type:         "beanies",
				Name:         "STAR EARTH",
				Color:        []string{"green"},
				Price:        80,
				Manufacturer: "ippal"},
			Stock: "INSTOCK",
		},
		Product{
			ProductWithoutStock: ProductWithoutStock{
				Id:           "BadID123",
				Type:         "beanies",
				Name:         "STAR EARTH",
				Color:        []string{"green"},
				Price:        80,
				Manufacturer: "jeppal"},
			Stock: "",
		},
	}
	return correctProducts
}

func getCorrectAvailabilityMap() map[string]string {
	correctAvailabilityMap := map[string]string{
		"ae8c8ad79a3e4a554d6f2":    "OUTOFSTOCK",
		"3af7caee9be9365e49e93576": "INSTOCK",
		"testerid123":              "LESSTHAN10",
	}
	return correctAvailabilityMap
}
func TestCreateAvailabilityMapCreatesCorrectMap(t *testing.T) {
	want := getCorrectAvailabilityMap()
	availabilityResponse := mock_service.MockAvailabilityResponse().Response
	got := createAvailabilityMap(availabilityResponse)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestAvailability(t *testing.T) {
	productResponse := mock_service.MockProductResponse().Response
	want := getCorrectProducts()
	availabilityMap := getCorrectAvailabilityMap()
	got := createProductsWithAvailability(availabilityMap, productResponse)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestProductSorting(t *testing.T) {
	var got = []Product{
		Product{
			ProductWithoutStock: ProductWithoutStock{
				Id:           "3af7caee9be9365e49e93576",
				Type:         "beanies",
				Name:         "STAR EARTH",
				Color:        []string{"green"},
				Price:        80,
				Manufacturer: "ippal"},
			Stock: "INSTOCK",
		},
		Product{
			ProductWithoutStock: ProductWithoutStock{
				Id:           "ae8c8ad79a3e4a554d6f2",
				Type:         "beanies",
				Name:         "SOPREV STAR",
				Color:        []string{"purple"},
				Price:        55,
				Manufacturer: "umpante",
			},
			Stock: "OUTOFSTOCK",
		},
		Product{
			ProductWithoutStock: ProductWithoutStock{
				Id:           "BadID123",
				Type:         "beanies",
				Name:         "STAR EARTH",
				Color:        []string{"green"},
				Price:        80,
				Manufacturer: "jeppal"},
			Stock: "",
		},
	}

	sort.Sort(ByName(got))
	want := getCorrectProducts()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}
