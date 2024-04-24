package udp_server

import (
	"broker/internal/cmd"
	"broker/internal/storage/responses"
	"broker/internal/storage/sensors"
	"fmt"
	"log"
	"strconv"
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
			fmt.Printf("O sensor (%s - %s) não foi encontrado\n", addr, cmd.IdFrom)
			return
		}

		if responses.FindBySensorId(cmd.IdFrom).SensorID == "" {
			responses.Create(cmd.IdFrom, cmd.Command)
		}

		value, err := strconv.ParseFloat(cmd.Content, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		responses.AddContent(cmd.IdFrom, value)
	})

	err := udpServer.Listen()
	if err != nil {
		log.Fatal("Error starting UDP server:", err)
	}
}
