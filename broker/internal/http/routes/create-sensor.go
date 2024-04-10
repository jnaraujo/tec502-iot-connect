package routes

import (
	"broker/internal/cmd"
	"broker/internal/sensor_conn"
	"broker/internal/storage"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type NewSensor struct {
	Address string `json:"address" validate:"required"`
	Id      string `json:"id" validate:"required"`
}

func CreateSensorHandler(c *gin.Context) {
	var newSensor NewSensor

	if err := c.BindJSON(&newSensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Corpo da requisição é inválido",
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(newSensor); err != nil {
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
