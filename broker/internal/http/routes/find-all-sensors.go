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
	sensors := sensors.FindSensors() // Busca todos os sensores

	sensorsWithStatus := []sensorWithOnlineStatus{}
	for _, sensor := range sensors {
		resp := responses.FindBySensorId(sensor.Id)
		isOnline := true

		// Verifica se o tempo existe e se o tempo é maior que o tempo máximo para ser considerado online
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
