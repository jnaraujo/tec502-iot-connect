package routes

import (
	"broker/errors"
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
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)

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

	if storage.GetSensorStorage().FindSensorNameByAddress(newSensor.Address) != "" {
		w.WriteHeader(http.StatusConflict)
		resp["message"] = "Sensor already registered"
		json.NewEncoder(w).Encode(resp)
		return
	}

	_, err = sensor.NewSensorConn(newSensor.Address)

	if err != nil {
		switch {
		case err == errors.ErrTimeout:
			w.WriteHeader(http.StatusServiceUnavailable)
			resp["message"] = "Time exceeded while trying to connect to sensor"
		case err == errors.ErrValidationFailed:
			w.WriteHeader(http.StatusUnauthorized)
			resp["message"] = "Sensor validation failed"
		default:
			w.WriteHeader(http.StatusInternalServerError)
			resp["message"] = "Internal server error"
		}

		json.NewEncoder(w).Encode(resp)
		return
	}

	storage.GetSensorStorage().AddSensor(newSensor.Name, newSensor.Address)

	w.WriteHeader(http.StatusCreated)

	resp["message"] = "Sensor registered"
	json.NewEncoder(w).Encode(resp)
}
