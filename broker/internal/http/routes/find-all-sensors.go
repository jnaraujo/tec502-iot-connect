package routes

import (
	"broker/internal/constants"
	"broker/internal/storage/responses"
	"broker/internal/storage/sensors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
			time.Since(resp.UpdatedAt.Time) > constants.MaxTimeToBeConsideredOnline {
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
