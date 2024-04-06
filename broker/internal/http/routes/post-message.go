package routes

import (
	"broker/internal/cmd_parser"
	"broker/internal/sensor_conn"
	"broker/internal/storage"
	"fmt"
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

	sensorData := storage.GetSensorDataStorage().Create(command.SensorID, command.Command, command.Content)

	_, err = sensor_conn.Request(addr, cmd_parser.Cmd{
		ID:      fmt.Sprintf("%d", sensorData.ID),
		Command: command.Command,
		Content: command.Content,
	})
	if err != nil {
		fmt.Println(err)
		storage.GetSensorDataStorage().UpdateResponse(sensorData.ID, "Sensor offline")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao enviar a mensagem para o sensor.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sensor":  command.SensorID,
		"command": command.Command,
		"content": command.Content,
	})
}
