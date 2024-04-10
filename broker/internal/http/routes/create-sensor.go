package routes

import (
	"broker/internal/errors"
	"broker/internal/sensorconn"
	"broker/internal/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewSensor struct {
	Address string `json:"address"`
	Id      string `json:"id"`
}

func CreateSensorHandler(c *gin.Context) {
	var newSensor NewSensor

	err := c.BindJSON(&newSensor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Corpo da requisição é inválido",
		})
		return
	}
	if newSensor.Address == "" || newSensor.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Corpo da requisição é inválido",
		})
		return
	}

	if storage.GetSensorStorage().DoesSensorExists(newSensor.Id, newSensor.Address) {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Sensor já existe",
		})
		return
	}

	_, err = sensorconn.New(newSensor.Address)
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

	storage.GetSensorStorage().AddSensor(newSensor.Id, newSensor.Address)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Sensor criado",
	})
}
