package mock_service

import (
	. "github.com/uberballo/webstore/model"
)

func MockProducts() ProductResponse {
	return ProductResponse{
		[]Product{
			Product{
				Id:           "ae8c8ad79a3e4a554d6f2",
				Type:         "beanies",
				Name:         "SOPREV STAR",
				Color:        []string{"purple"},
				Price:        55,
				Manufacturer: "umpante",
			},
			Product{
				Id:           "3af7caee9be9365e49e93576",
				Type:         "beanies",
				Name:         "STAR EARTH",
				Color:        []string{"green"},
				Price:        80,
				Manufacturer: "ippal",
			},
			Product{
				Id:           "BadID123",
				Type:         "beanies",
				Name:         "STAR EARTH",
				Color:        []string{"green"},
				Price:        80,
				Manufacturer: "jeppal",
			},
		},
	}
}
