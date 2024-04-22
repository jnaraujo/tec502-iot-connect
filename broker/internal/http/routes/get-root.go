package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Broker API",
	})
}
