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
	resp := make(map[string]string)

	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body")
		return
	}

	if message.SensorID == "" || message.Command == "" {
		w.WriteHeader(http.StatusBadRequest)
		resp["message"] = "Invalid request body"
		json.NewEncoder(w).Encode(resp)
		return
	}

	addr := storage.GetSensorStorage().FindSensorAddrByName(message.SensorID)

	if addr == "" {
		w.WriteHeader(http.StatusNotFound)
		resp["message"] = "Sensor not found"
		json.NewEncoder(w).Encode(resp)
		return
	}

	conn, err := sensor.NewSensorConn(addr)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		resp["message"] = "Error sending message to sensor"
		json.NewEncoder(w).Encode(resp)
		return
	}

	defer conn.Close()

	response, err := conn.Request(message.Command, message.Content)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		resp["message"] = "Error sending message to sensor"
		json.NewEncoder(w).Encode(resp)
		return
	}

	sensorResArr := strings.Split(response, "\n\n")

	w.WriteHeader(http.StatusOK)
	resp["command"] = strings.Split(sensorResArr[0], ": ")[1]
	resp["content"] = sensorResArr[1]
	json.NewEncoder(w).Encode(resp)
}
