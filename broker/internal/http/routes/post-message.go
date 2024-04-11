package routes

import (
	"broker/internal/cmd"
	"broker/internal/sensor_conn"
	"broker/internal/storage/sensors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type CommandRequest struct {
	SensorID string `json:"sensor_id" validate:"required"`
	Command  string `json:"command" validate:"required"`
	Content  string `json:"content"`
}

func PostMessageHandler(c *gin.Context) {
	var command CommandRequest
	if err := c.BindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Corpo da requisição é inválido",
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(command); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Corpo da requisição é inválido",
		})
		return
	}

	addr := sensors.FindSensorAddrById(command.SensorID)
	if addr == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Sensor not found",
		})
		return
	}

	response, err := sensor_conn.Request(addr, cmd.New(
		"BROKER", command.SensorID, command.Command, command.Content,
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
