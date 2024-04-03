package routes

import (
	"encoding/json"
	"net/http"
)

func GetRootHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)

	resp["message"] = "Hello, world!"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
