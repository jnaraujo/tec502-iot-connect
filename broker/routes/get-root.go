package routes

import (
	"encoding/json"
	"net/http"
)

func GetRootHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to the Broker API! 😘🤞"})
}
