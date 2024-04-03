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
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body")
		return
	}

	if message.SensorID == "" || message.Content == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body")
		return
	}

	fmt.Printf("Received message from sensor %s: %s\n", message.SensorID, message.Content)

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "Message received")
}
