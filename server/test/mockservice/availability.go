package mockservice

import (
	. "github.com/uberballo/warehouse/server/model"
)

//MockAvailabilityResponse returns a correct availabilityResponse mock
func MockAvailabilityResponse() AvailabilityResponse {
	return AvailabilityResponse{
		Response: []Availability{
			{
				Id:          "3AF7CAEE9BE9365E49E93576",
				Datapayload: "<AVAILABILITY> <CODE>200</CODE> <INSTOCKVALUE>INSTOCK</INSTOCKVALUE> </AVAILABILITY>",
			},
			{
				Id:          "AE8C8AD79A3E4A554D6F2",
				Datapayload: "<AVAILABILITY> <CODE>200</CODE> <INSTOCKVALUE>OUTOFSTOCK</INSTOCKVALUE> </AVAILABILITY>",
			},
			{
				Id:          "testerid123",
				Datapayload: "<AVAILABILITY> <CODE>200</CODE> <INSTOCKVALUE>LESSTHAN10</INSTOCKVALUE> </AVAILABILITY>",
			},
		},
	}
}
