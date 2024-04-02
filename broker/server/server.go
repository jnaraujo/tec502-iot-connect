package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	DefaultPort = 8080
)

func Init() {
	r := mux.NewRouter()

	r.HandleFunc("/", getRootHandler).Methods("GET")
	r.HandleFunc("/message", postMessageHandler).Methods("POST")

	fmt.Println("Server started on port", DefaultPort)

	err := http.ListenAndServe(fmt.Sprintf(":%d", DefaultPort), r)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func getRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

type Message struct {
	SensorID string `json:"sensor_id"`
	Content  string `json:"content"`
}

func postMessageHandler(w http.ResponseWriter, r *http.Request) {

	// parse request body
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
