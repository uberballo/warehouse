package model

type Availability struct {
	Id          string
	Datapayload string
}

type AvailabilityResponse struct {
	Response []Availability
}

type Product struct {
	Id           string
	Type         string
	Name         string
	Color        []string
	Price        int
	Manufacturer string
}

type ProductResponse struct {
	Response []Product
}
