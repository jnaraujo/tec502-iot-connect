package routes

import (
	"broker/internal/storage/responses"
	"broker/internal/storage/sensors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteSensorRoute(c *gin.Context) {
	sensorId := c.Param("sensor_id")

	if sensors.FindSensorAddrById(sensorId) == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Sensor não encontrado.",
		})
		return
	}

	responses.DeleteBySensorId(sensorId)

	sensors.DeleteSensorBySensorId(sensorId)
	c.JSON(http.StatusOK, gin.H{
		"message": "Sensor deletado.",
	})
}
