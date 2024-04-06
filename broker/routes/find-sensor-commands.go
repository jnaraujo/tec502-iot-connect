package routes

import (
	"broker/cmd_parser"
	"broker/sensor"
	"broker/storage"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func FindSensorCommands(c *gin.Context) {
	sensor_id := c.Param("sensor_id")

	addr := storage.GetSensorStorage().FindSensorAddrByName(sensor_id)

	if addr == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Sensor não encontrado.",
		})
		return
	}

	conn, err := sensor.NewSensorConn(addr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error connecting to sensor",
		})
		return
	}
	defer conn.Close()

	resp, err := conn.Request(cmd_parser.Cmd{
		ID:      "#",
		Command: "get_commands",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao realizar a requisição com o sensor.",
		})
		return
	}

	cmd, err := cmd_parser.DecodeCmd(resp)
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
