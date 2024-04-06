package routes

import (
	"broker/errors"
	"broker/sensor"
	"broker/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewSensor struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

func CreateSensorHandler(c *gin.Context) {
	var newSensor NewSensor

	err := c.BindJSON(&newSensor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	if newSensor.Address == "" || newSensor.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	if storage.GetSensorStorage().FindSensorIdByAddress(newSensor.Address) != "" {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Sensor already exists",
		})
		return
	}

	_, err = sensor.NewSensorConn(newSensor.Address)

	if err != nil {
		switch {
		case err == errors.ErrTimeout:
			c.JSON(http.StatusRequestTimeout, gin.H{
				"message": "Sensor connection timeout",
			})
		case err == errors.ErrValidationFailed:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid sensor address",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error connecting to sensor",
			})
		}
		return
	}

	storage.GetSensorStorage().AddSensor(newSensor.Name, newSensor.Address)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Sensor created",
	})
}
