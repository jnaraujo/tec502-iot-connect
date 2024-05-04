package http

import (
	"broker/internal/http/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Cria um novo servidor HTTP
func NewServer(addr string, port int) {
	fmt.Printf("Server started on %s:%d\n", addr, port)

	g := gin.New()

	g.Use(gin.Recovery())
	g.Use(corsMiddleware())

	RegisterRoutes(g)

	err := g.Run(fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// Registra as rotas do servidor HTTP
func RegisterRoutes(g *gin.Engine) {
	g.GET("/", routes.GetRootHandler)
	g.POST("/message", routes.PostMessageHandler)
	g.POST("/sensor", routes.CreateSensorHandler)
	g.GET("/sensor", routes.FindAllSensorsHandler)
	g.GET("/sensor/commands/:sensor_id", routes.FindSensorCommands)
	g.GET("/sensor/data", routes.FindAllSensorDataHandler)
	g.DELETE("/sensor/:sensor_id", routes.DeleteSensorRoute)
}

// Middleware para permitir requisições de outros domínios
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
