package http

import (
	"broker/internal/http/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewServer(addr string, port int) {
	g := gin.Default()
	g.Use(corsMiddleware())

	registerRoutes(g)

	fmt.Printf("Server started on %s:%d\n", addr, port)

	err := g.Run(fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func registerRoutes(g *gin.Engine) {
	g.GET("/", routes.GetRootHandler)
	g.POST("/message", routes.PostMessageHandler)
	g.POST("/sensor", routes.CreateSensorHandler)
	g.GET("/sensor", routes.FindAllSensorsHandler)
	g.GET("/sensor/commands/:sensor_id", routes.FindSensorCommands)
	g.GET("/sensor/data", routes.FindAllSensorDataHandler)
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
