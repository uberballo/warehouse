package model

type Availability struct {
	Id          string
	Datapayload string
}

type AvailabilityResponse struct {
	Response []Availability
}

type ProductWithoutStock struct {
	Id           string
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
