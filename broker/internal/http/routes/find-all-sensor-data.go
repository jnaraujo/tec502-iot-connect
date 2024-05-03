package routes

import (
	"broker/internal/storage/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAllSensorDataHandler(c *gin.Context) {
	data := responses.FindAll() // Busca todas as respostas dos sensores

	c.JSON(http.StatusOK, data)
}
