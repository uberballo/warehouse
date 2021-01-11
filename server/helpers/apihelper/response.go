package apihelper

import (
	"encoding/json"
	"net/http"
)

//Respond creates json response
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
