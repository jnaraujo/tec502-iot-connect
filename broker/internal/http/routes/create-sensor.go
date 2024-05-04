package routes

import (
	"broker/internal/cmd"
	"broker/internal/sensor_conn"
	"broker/internal/storage/sensors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type CreateSensorBody struct {
	Address string `json:"address" binding:"required"`
	Id      string `json:"id" binding:"required"`
}

func CreateSensorHandler(c *gin.Context) {
	var body CreateSensorBody
	// O método ShouldBindJSON é responsável transformar o corpo da requisição em um objeto e validar se o corpo da requisição está de acordo com o esperado.
	if err := c.ShouldBindJSON(&body); err != nil {
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

	// O método Validate é responsável por verificar se a conexão com o sensor é válida, ou seja, se o sensor está respondendo.
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
