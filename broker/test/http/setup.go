package http

import (
	"broker/internal/http"
	"broker/internal/sensor_conn"
	"broker/internal/storage"
	"broker/internal/udp_server"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(gin.Recovery())
	http.RegisterRoutes(r)
	return r
}

func setupUdpServer() *udp_server.Server {
	server := udp_server.NewServer("0.0.0.0", 5310)
	return server
}

func checkSensor() {
	// Verifica se o sensor esta online
	err := sensor_conn.Validate("localhost:3399")
	if err != nil {
		panic("O *Sensor do Ar Condicionado* precisa estar online para o teste! Se ele já estiver ligado, verifica se está ouvindo na porta '3399'.")
	}
}

func cleanUp() {
	// limpas os storages
	storage.ClearAll()
}
