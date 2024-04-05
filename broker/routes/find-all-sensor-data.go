package routes

import (
	"broker/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAllSensorDataHandler(c *gin.Context) {
	data := storage.GetSensorDataStorage().FindAll()

	c.JSON(http.StatusOK, data)
}
