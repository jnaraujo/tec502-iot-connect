package routes

import (
	"broker/internal/cmd"
	"broker/internal/sensor_conn"
	"broker/internal/storage/sensors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type PostMessageBody struct {
	SensorID string `json:"sensor_id" validate:"required"`
	Command  string `json:"command" validate:"required"`
	Content  string `json:"content"`
}

func PostMessageHandler(c *gin.Context) {
	var body PostMessageBody
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

	addr := sensors.FindSensorAddrById(body.SensorID)
	if addr == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Sensor not found",
		})
		return
	}

	response, err := sensor_conn.Request(addr, cmd.New(
		"BROKER", body.SensorID, body.Command, body.Content,
	))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao enviar a mensagem para o sensor.",
		})
		return
	}

	cmd, err := cmd.Encode(response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro decodificar mensagem do sensor.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": cmd.Content,
	})
}
