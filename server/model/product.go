package model

//CombineProductResponses combines slice of multiple ProductResponse to single ProductResponse
func CombineProductResponses(productResponse []ProductResponse) ProductResponse {
	responses := []ProductWithoutStock{}
	for _, pr := range productResponse {
		responses = append(responses, pr.Response...)
	}
	return ProductResponse{Response: responses}
}

//CombineAvailabilityResponses combines slice of multiple AvailabilityResponses into single AvailabilityResponse
func CombineAvailabilityResponses(availabilityResponse []AvailabilityResponse) AvailabilityResponse {
	responses := []Availability{}
	for _, pr := range availabilityResponse {
		responses = append(responses, pr.Response...)
	}
	return AvailabilityResponse{Response: responses}
}

func (p ByName) Len() int           { return len(p) }
func (p ByName) Less(i, j int) bool { return p[i].Name < p[j].Name }
func (p ByName) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
