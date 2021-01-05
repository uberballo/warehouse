package internal

import (
	"reflect"
	"testing"

	"github.com/uberballo/webstore/test/mock_service"

	. "github.com/uberballo/webstore/model"
)

func getCorrectProducts() []Product {
	var correctProducts = []Product{
		Product{
			Id:           "ae8c8ad79a3e4a554d6f2",
			Type:         "beanies",
			Name:         "SOPREV STAR",
			Color:        []string{"purple"},
			Price:        55,
			Manufacturer: "umpante",
			Stock:        "INSTOCK",
		},
		Product{
			Id:           "3af7caee9be9365e49e93576",
			Type:         "beanies",
			Name:         "STAR EARTH",
			Color:        []string{"green"},
			Price:        80,
			Manufacturer: "ippal",
			Stock:        "OUTOFSTOCK",
		},
		Product{
			Id:           "BadID123",
			Type:         "beanies",
			Name:         "STAR EARTH",
			Color:        []string{"green"},
			Price:        80,
			Manufacturer: "jeppal",
			Stock:        "",
		},
	}
	return correctProducts
}

var correctAvailabilityMap = map[string]string{
	"ae8c8ad79a3e4a554d6f2":    "INSTOCK",
	"3af7caee9be9365e49e93576": "OUTOFSTOCK",
	"testerid123":              "LESSTHAN10",
}

func TestAvailability(t *testing.T) {
	productResponse := mock_service.MockProducts().Response
	want := getCorrectProducts()
	got := ReadInStockValue(correctAvailabilityMap, productResponse)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
