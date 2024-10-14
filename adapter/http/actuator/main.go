package actuator

import (
	"encoding/json"
	"net/http"
)

// HealthBody is used to return the status of the server
type HealthBody struct {
	Status string `json:"status"`
}

// Health is used to check if the server is alive
func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var status = HealthBody{"Alive"}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(status)
}