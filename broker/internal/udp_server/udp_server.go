package udp_server

import (
	"broker/internal/cmd_parser"
	"broker/internal/storage"
	"fmt"
	"log"
	"strconv"
)

func NewServer(addr string, port int) {
	fmt.Printf("Starting UDP server on %s:%d", addr, port)

	udpServer := NewUDPServer(fmt.Sprintf("%s:%d", addr, port))
	defer udpServer.Close()

	udpServer.HandleRequest(func(msg string, reply func(string) error) {
		cmd, err := cmd_parser.DecodeCmd(msg)
		if err != nil {
			reply("Invalid command")
			return
		}

		numId, err := strconv.Atoi(cmd.ID)
		if err != nil {
			reply("Invalid sensor ID")
			return
		}

		data := storage.GetSensorDataStorage().FindByID(numId)
		if data == nil {
			reply("Sensor not found")
			return
		}

		storage.GetSensorDataStorage().UpdateResponse(data.ID, cmd.Content)

		fmt.Println("Received response from request", data.ID)
		reply("Message received")
	})

	err := udpServer.Listen()
	if err != nil {
		log.Fatal("Error starting UDP server:", err)
	}
}
