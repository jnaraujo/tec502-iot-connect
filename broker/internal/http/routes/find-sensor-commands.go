package routes

import (
	"broker/internal/cmd"
	"broker/internal/sensor_conn"
	"broker/internal/storage/sensors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func FindSensorCommands(c *gin.Context) {
	sensor_id := c.Param("sensor_id") // /sensors/commands/:sensor_id

	// Verifica se o sensor existe
	addr := sensors.FindSensorAddrById(sensor_id)
	if addr == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Sensor não encontrado.",
		})
		return
	}

	// Envia o comando para o sensor para obter os comandos disponíveis
	resp, err := sensor_conn.Request(addr, cmd.New("BROKER", sensor_id, "get_commands", ""))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao realizar a requisição com o sensor.",
		})
		return
	}

	cmd, err := cmd.Encode(resp)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "O comando recebido do sensor é inválido.",
		})
		return
	}

	// Separa os comandos
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
