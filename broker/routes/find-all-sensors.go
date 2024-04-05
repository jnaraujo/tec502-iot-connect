package routes

import (
	"broker/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAllSensorsHandler(c *gin.Context) {
	sensors := storage.GetSensorStorage().GetSensors()

	c.JSON(http.StatusOK, gin.H{
		"sensors": sensors,
	})
}
