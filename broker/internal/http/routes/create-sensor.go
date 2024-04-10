package routes

import (
	"broker/internal/cmd"
	"broker/internal/sensor_conn"
	"broker/internal/storage"
	"net/http"
	"os"

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

	conn, err := sensor_conn.New(newSensor.Address)
	if err != nil {
		switch {
		case os.IsTimeout(err):
			c.JSON(http.StatusRequestTimeout, gin.H{
				"message": "O sensor demorou demais para responder",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao se conectar com o sensor",
			})
		}
		return
	}

	_, err = conn.Send(cmd.New("BROKER", newSensor.Id, "set_id", newSensor.Id).Decode())
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Erro ao setar o id no sensor.",
		})
		return
	}

	storage.GetSensorStorage().AddSensor(newSensor.Id, newSensor.Address)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Sensor criado",
	})
}
