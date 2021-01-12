package model

//Availability of a certain product
type Availability struct {
	ID          string //Product ID
	Datapayload string
}

type AvailabilityResponse struct {
	Response []Availability
}

type ProductWithoutStock struct {
	ID           string
	Type         string
	Name         string
	Color        []string
	Price        int
	Manufacturer string
}

type Product struct {
	Stock string
	ProductWithoutStock
}

type ProductResponse struct {
	Category string
	Response []ProductWithoutStock
}

type BadAPiResponse struct {
	ProductResponses      []ProductResponse
	AvailabilityResponses []AvailabilityResponse
	AvailabilityErrors    []string
}

type ByName []Product
