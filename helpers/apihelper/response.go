package apihelper

import (
	"encoding/json"
	"net/http"
)

type ResponseData struct {
	Data interface{} `json:"data"`
}

func Respond(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
