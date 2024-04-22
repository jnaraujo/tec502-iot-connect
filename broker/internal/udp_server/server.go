package udp_server

import (
	"broker/internal/cmd"
	"broker/internal/storage/responses"
	"broker/internal/storage/sensors"
	"fmt"
	"log"
)

func NewServer(addr string, port int) {
	fmt.Printf("Starting UDP server on %s:%d", addr, port)

	udpServer := NewUDPServer(fmt.Sprintf("%s:%d", addr, port))
	defer udpServer.Close()

	udpServer.HandleRequest(func(addr, content string) {
		cmd, err := cmd.Encode(content)
		if err != nil {
			fmt.Println("Erro ao encodar o comando", err)
			return
		}

		if sensors.FindSensorAddrById(cmd.IdFrom) == "" {
			fmt.Println("O sensor não foi encontrado")
			return
		}

		response := responses.FindBySensorId(cmd.IdFrom)
		if response.SensorID == "" {
			responses.Create(cmd.IdFrom, cmd.Command, cmd.Content)
			return
		}

		responses.UpdateContent(cmd.IdFrom, cmd.Content)
	})

	err := udpServer.Listen()
	if err != nil {
		log.Fatal("Error starting UDP server:", err)
	}
}
