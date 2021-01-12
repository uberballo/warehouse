package mockservice

import (
	. "github.com/uberballo/warehouse/server/model"
)

//MockProductResponse returns a correct ProductResponse mock
func MockProductResponse() ProductResponse {
	return ProductResponse{
		Category: "beanies",
		Response: []ProductWithoutStock{
			{
				ID:           "ae8c8ad79a3e4a554d6f2",
				Type:         "beanies",
				Name:         "SOPREV STAR",
				Color:        []string{"purple"},
				Price:        55,
				Manufacturer: "umpante",
			},
			{
				ID:           "3af7caee9be9365e49e93576",
				Type:         "beanies",
				Name:         "STAR EARTH",
				Color:        []string{"green"},
				Price:        80,
				Manufacturer: "ippal",
			},
			{
				ID:           "BadID123",
				Type:         "beanies",
				Name:         "STAR EARTH",
				Color:        []string{"green"},
				Price:        80,
				Manufacturer: "jeppal",
			},
		},
	}
}
