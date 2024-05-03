package routes

import (
	"broker/internal/storage/responses"
	"broker/internal/storage/sensors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteSensorRoute(c *gin.Context) {
	sensorId := c.Param("sensor_id")

	// Verifica se o sensor existe
	if sensors.FindSensorAddrById(sensorId) == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Sensor não encontrado.",
		})
		return
	}

	// Deleta todas as respostas do sensor
	responses.DeleteBySensorId(sensorId)

	// Deleta o sensor
	sensors.DeleteSensorBySensorId(sensorId)

	c.JSON(http.StatusOK, gin.H{
		"message": "Sensor deletado.",
	})
}
