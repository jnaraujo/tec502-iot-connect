package routes

import (
	"broker/sensor"
	"broker/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

type CommandRequest struct {
	SensorID string `json:"sensor_id"`
	Command  string `json:"command"`
	Content  string `json:"content"`
}

func PostMessageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var command CommandRequest
	err := json.NewDecoder(r.Body).Decode(&command)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body")
		return
	}

	if command.SensorID == "" || command.Command == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid request body",
		})
		return
	}

	addr := storage.GetSensorStorage().FindSensorAddrByName(command.SensorID)

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

	_, err = conn.Request(command.Command, command.Content)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Error sending message to sensor",
		})
		return
	}

	storage.GetSensorDataStorage().Create(command.SensorID, command.Command, command.Content)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Message sent",
	})
}
