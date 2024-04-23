package routes

import (
	"broker/internal/cmd"
	"broker/internal/sensor_conn"
	"broker/internal/storage/sensors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type CreateSensorBody struct {
	Address string `json:"address" validate:"required"`
	Id      string `json:"id" validate:"required"`
}

func CreateSensorHandler(c *gin.Context) {
	var body CreateSensorBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Corpo da requisição é inválido",
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Corpo da requisição é inválido",
		})
		return
	}

	if sensors.DoesSensorExists(body.Id, body.Address) {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Sensor já existe",
		})
		return
	}

	err := sensor_conn.Validate(body.Address)
	if err != nil {
		fmt.Println(err)
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

	_, err = sensor_conn.Request(
		body.Address,
		cmd.New("BROKER", body.Id, "set_id", body.Id),
	)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Erro ao definir o id no sensor.",
		})
		return
	}

	sensors.AddSensor(body.Id, body.Address)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Sensor criado",
	})
}
