package apihelper

import (
	"net/http"
	"time"
)

//GetHTTPClient creates a HTTP client with timeout
func GetHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 20,
	}
}
