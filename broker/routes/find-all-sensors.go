package routes

import (
	"broker/storage"
	"encoding/json"
	"net/http"
)

func FindAllSensorsHandler(w http.ResponseWriter, r *http.Request) {
	sensors := storage.GetSensorStorage().GetSensors()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sensors)
}
