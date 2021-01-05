package mock_service

import (
	. "github.com/uberballo/webstore/model"
)

func MockAvailability() AvailabilityResponse {
	return AvailabilityResponse{
		[]Availability{
			Availability{
				Id:          "3AF7CAEE9BE9365E49E93576",
				Datapayload: "<AVAILABILITY> <CODE>200</CODE> <INSTOCKVALUE>INSTOCK</INSTOCKVALUE> </AVAILABILITY>",
			},
			Availability{
				Id:          "AE8C8AD79A3E4A554D6F2",
				Datapayload: "<AVAILABILITY> <CODE>200</CODE> <INSTOCKVALUE>OUTOFSTOCK</INSTOCKVALUE> </AVAILABILITY>",
			},
			Availability{
				Id:          "None",
				Datapayload: "<AVAILABILITY> <CODE>200</CODE> <INSTOCKVALUE>INSTOCK</INSTOCKVALUE> </AVAILABILITY>",
			},
		},
	}
}
