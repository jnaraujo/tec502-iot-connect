package routes

import (
	"broker/sensor"
	"broker/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Message struct {
	SensorID string `json:"sensor_id"`
	Command  string `json:"command"`
	Content  string `json:"content"`
}

func PostMessageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body")
		return
	}

	if message.SensorID == "" || message.Command == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid request body",
		})
		return
	}

	addr := storage.GetSensorStorage().FindSensorAddrByName(message.SensorID)

	if addr == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Sensor not found",
		})
		return
	}

	conn, err := sensor.NewSensorConn(addr)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Error sending message to sensor",
		})
		return
	}

	defer conn.Close()

	response, err := conn.Request(message.Command, message.Content)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Error sending message to sensor",
		})
		return
	}

	sensorResArr := strings.Split(response, "\n\n")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":  strings.Split(sensorResArr[0], ": ")[1],
		"response": sensorResArr[1],
	})
}
