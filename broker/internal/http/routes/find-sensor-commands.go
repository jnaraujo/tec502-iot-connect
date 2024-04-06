package routes

import (
	"broker/internal/cmd"
	"broker/internal/sensorconn"
	"broker/internal/storage"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func FindSensorCommands(c *gin.Context) {
	sensor_id := c.Param("sensor_id")

	addr := storage.GetSensorStorage().FindSensorAddrById(sensor_id)

	if addr == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Sensor não encontrado.",
		})
		return
	}

	resp, err := sensorconn.Request(addr, cmd.New("#", "get_commands", ""))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao realizar a requisição com o sensor.",
		})
		return
	}

	cmd, err := cmd.Encode(resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "O comando recebido do sensor é inválido.",
		})
		return
	}

	commands := strings.Split(cmd.Content, ", ")

	if commands == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "o sensor não existe",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"commands": commands,
	})
}
