package routes

import (
	"broker/internal/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAllSensorDataHandler(c *gin.Context) {
	data := storage.GetSensorResponseStorage().FindAll()

	c.JSON(http.StatusOK, data)
}
