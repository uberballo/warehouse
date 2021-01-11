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

func (p ByName) Len() int           { return len(p) }
func (p ByName) Less(i, j int) bool { return p[i].Name < p[j].Name }
func (p ByName) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
