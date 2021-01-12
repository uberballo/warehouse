package model

//Availability of a certain product
type Availability struct {
	ID          string //Product ID
	Datapayload string
}

//AvailabilityResponse contains list of Availability structs
type AvailabilityResponse struct {
	Response []Availability
}

//ProductWithoutStock represents single product without stock
type ProductWithoutStock struct {
	ID           string
	Type         string
	Name         string
	Color        []string
	Price        int
	Manufacturer string
}

//Product contains ProductWithoutStock and availability amount
type Product struct {
	Stock string
	ProductWithoutStock
}

//ProductResponse represents response from badapi/products/:category
type ProductResponse struct {
	Category string
	Response []ProductWithoutStock
}

//BadAPIResponse represents combined responses from badapi's Products and Availability
type BadAPIResponse struct {
	ProductResponses      []ProductResponse
	AvailabilityResponses []AvailabilityResponse
	AvailabilityErrors    []string
}

//ByName sorting order struct for Products
type ByName []Product
