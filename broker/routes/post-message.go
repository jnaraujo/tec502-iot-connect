package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	SensorID string `json:"sensor_id"`
	Content  string `json:"content"`
}

func PostMessageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)

	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body")
		return
	}

	if message.SensorID == "" || message.Content == "" {
		w.WriteHeader(http.StatusBadRequest)
		resp["message"] = "Invalid request body"
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusAccepted)

	resp["message"] = "Message received"
	json.NewEncoder(w).Encode(resp)
}
