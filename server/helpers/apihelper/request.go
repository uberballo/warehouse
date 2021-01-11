package apihelper

import (
	"net/http"
	"time"
)

func GetHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 20,
	}
}
