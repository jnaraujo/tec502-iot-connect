package routes

import (
	"broker/storage"
	"encoding/json"
	"net/http"
)

func FindAllSensorDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := storage.GetSensorDataStorage().FindAll()

	json.NewEncoder(w).Encode(data)
}
