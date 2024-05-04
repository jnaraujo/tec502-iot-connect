package routes

import (
	"broker/internal/cmd"
	"broker/internal/sensor_conn"
	"broker/internal/storage/sensors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostMessageBody struct {
	SensorID string `json:"sensor_id" binding:"required"`
	Command  string `json:"command" binding:"required"`
	Content  string `json:"content"`
}

func PostMessageHandler(c *gin.Context) {
	var body PostMessageBody
	// O método ShouldBindJSON é responsável transformar o corpo da requisição em um objeto e validar se o corpo da requisição está de acordo com o esperado.
	if err := c.ShouldBindJSON(&body); err != nil {
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

	// Envia o comando para o sensor e aguarda a resposta
	response, err := sensor_conn.Request(addr, cmd.New(
		"BROKER", body.SensorID, body.Command, body.Content,
	))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao enviar a mensagem para o sensor.",
		})
		return
	}

	// Decodifica a resposta do sensor
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
