package routes

import (
	"broker/sensor"
	"broker/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

type NewSensor struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

func CreateSensorHandler(w http.ResponseWriter, r *http.Request) {
	var newSensor NewSensor

	err := json.NewDecoder(r.Body).Decode(&newSensor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body")
		return
	}

	if newSensor.Address == "" || newSensor.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body")
		return
	}

	if storage.GetSensorStorage().FindSensorByAddress(newSensor.Address) != nil {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "Sensor already registered")
		return
	}

	sensor, err := sensor.NewSensorConn(sensor.NewSensor{
		Name:    newSensor.Name,
		Address: newSensor.Address,
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid sensor address")
		return
	}

	storage.GetSensorStorage().AddSensor(*sensor)

	resp := make(map[string]string)
	resp["message"] = "Sensor registered"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
