package routes

import (
	"broker/internal/storage/responses"
	"broker/internal/storage/sensors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	maxTimeSinceLastResponse = 10 * time.Second
)

type sensorWithOnlineStatus struct {
	sensors.Sensor
	IsOnline bool `json:"is_online"`
}

func FindAllSensorsHandler(c *gin.Context) {
	sensors := sensors.FindSensors()

	sensorsWithStatus := []sensorWithOnlineStatus{}

	for _, sensor := range sensors {
		resp := responses.FindBySensorId(sensor.Id)
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
