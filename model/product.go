package model

func CombineProductResponses(productResponse []ProductResponse) ProductResponse {
	responses := []ProductWithoutStock{}
	for _, pr := range productResponse {
		responses = append(responses, pr.Response...)
	}
	return ProductResponse{Response: responses}
}

func CombinAvailabilityResponses(availabilityResponse []AvailabilityResponse) AvailabilityResponse {
	responses := []Availability{}
	for _, pr := range availabilityResponse {
		responses = append(responses, pr.Response...)
	}
	return AvailabilityResponse{Response: responses}
}
