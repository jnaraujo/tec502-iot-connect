package routes

import (
	"broker/internal/storage"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	maxTimeSinceLastResponse = 10 * time.Second
)

type sensorWithOnlineStatus struct {
	storage.Sensor
	IsOnline bool `json:"is_online"`
}

func FindAllSensorsHandler(c *gin.Context) {
	sensors := storage.GetSensorStorage().FindSensors()

	sensorsWithStatus := []sensorWithOnlineStatus{}

	for _, sensor := range sensors {
		resp := storage.GetSensorResponseStorage().FindBySensorId(sensor.Id)
		isOnline := true

		if resp.UpdatedAt.Time.IsZero() ||
			time.Since(resp.UpdatedAt.Time) > maxTimeSinceLastResponse {
			isOnline = false
		}

		sensorsWithStatus = append(sensorsWithStatus, sensorWithOnlineStatus{
			Sensor:   sensor,
			IsOnline: isOnline,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"sensors": sensorsWithStatus,
	})
}
