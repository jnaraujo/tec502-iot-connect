package routes

import (
	"broker/internal/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteSensorRoute(c *gin.Context) {
	sensorId := c.Param("sensor_id")

	if storage.GetSensorStorage().FindSensorAddrById(sensorId) == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Sensor não encontrado.",
		})
		return
	}

	storage.GetSensorStorage().DeleteSensorBySensorId(sensorId)
	c.JSON(http.StatusOK, gin.H{
		"message": "Sensor deletado.",
	})
}
