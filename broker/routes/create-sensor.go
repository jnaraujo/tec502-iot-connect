package routes

import (
	"broker/errors"
	"broker/sensor"
	"broker/storage"
	"encoding/json"
	"net/http"
)

type NewSensor struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

func CreateSensorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newSensor NewSensor

	err := json.NewDecoder(r.Body).Decode(&newSensor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid request body",
		})
		return
	}

	if newSensor.Address == "" || newSensor.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid request body",
		})
		return
	}

	if storage.GetSensorStorage().FindSensorNameByAddress(newSensor.Address) != "" {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Sensor already registered",
		})
		return
	}

	_, err = sensor.NewSensorConn(newSensor.Address)

	if err != nil {
		var msg string

		switch {
		case err == errors.ErrTimeout:
			w.WriteHeader(http.StatusServiceUnavailable)
			msg = "Time exceeded while trying to connect to sensor"
		case err == errors.ErrValidationFailed:
			w.WriteHeader(http.StatusUnauthorized)
			msg = "Sensor validation failed"
		default:
			w.WriteHeader(http.StatusInternalServerError)
			msg = "Error connecting to sensor"
		}

		json.NewEncoder(w).Encode(map[string]string{
			"message": msg,
		})
		return
	}

	storage.GetSensorStorage().AddSensor(newSensor.Name, newSensor.Address)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Sensor created",
	})
}
