package helpers

import (
	"reflect"
	"sort"
	"testing"

	. "github.com/uberballo/warehouse/server/model"
	"github.com/uberballo/warehouse/server/test/mockservice"
)

func getCorrectProducts() []Product {
	var correctProducts = []Product{
		{
			ProductWithoutStock: ProductWithoutStock{
				ID:           "ae8c8ad79a3e4a554d6f2",
				Type:         "beanies",
				Name:         "SOPREV STAR",
				Color:        []string{"purple"},
				Price:        55,
				Manufacturer: "umpante",
			},
			Stock: "OUTOFSTOCK",
		},
		{
			ProductWithoutStock: ProductWithoutStock{
				ID:           "3af7caee9be9365e49e93576",
				Type:         "beanies",
				Name:         "STAR EARTH",
				Color:        []string{"green"},
				Price:        80,
				Manufacturer: "ippal"},
			Stock: "INSTOCK",
		},
		{
			ProductWithoutStock: ProductWithoutStock{
				ID:           "BadID123",
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
	availabilityResponse := mockservice.MockAvailabilityResponse().Response
	got := createAvailabilityMap(availabilityResponse)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestAvailability(t *testing.T) {
	productResponse := mockservice.MockProductResponse().Response
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

func TestGetInStockValue(t *testing.T) {
	payload := "<AVAILABILITY> <CODE>200</CODE> <INSTOCKVALUE>OUTOFSTOCK</INSTOCKVALUE> </AVAILABILITY>"
	want := "OUTOFSTOCK"
	t.Run(payload, testGetInStockValueFunc(payload, want))
	t.Run("", testGetInStockValueFunc("nothing", "nothing"))
}

func testGetInStockValueFunc(payload, want string) func(*testing.T) {
	return func(t *testing.T) {
		got := getInStockValue(payload)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
