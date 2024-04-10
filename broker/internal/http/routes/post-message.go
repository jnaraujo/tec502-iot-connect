package routes

import (
	"broker/internal/cmd"
	"broker/internal/sensorconn"
	"broker/internal/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommandRequest struct {
	SensorID string `json:"sensor_id"`
	Command  string `json:"command"`
	Content  string `json:"content"`
}

func PostMessageHandler(c *gin.Context) {
	var command CommandRequest

	err := c.BindJSON(&command)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	if command.SensorID == "" || command.Command == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	addr := storage.GetSensorStorage().FindSensorAddrById(command.SensorID)
	if addr == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Sensor not found",
		})
		return
	}

	response, err := sensorconn.Request(addr, cmd.New(
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
		"response": cmd.Content,
	})
}
