package main

import (
	"broker/cmd_parser"
	"broker/routes"
	"broker/storage"
	"broker/udp"
	"broker/utils"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

const (
	defaultPort   = 8080
	udoServerPort = 5310
)

func main() {
	showWelcomeMsg()

	go handleUdpServer()
	handleServer()
}

func handleUdpServer() {
	fmt.Println("Starting UDP server on port", udoServerPort)

	udpServer := udp.NewUDPServer(fmt.Sprintf(":%d", udoServerPort))
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

func handleServer() {
	router := gin.Default()

	router.GET("/", routes.GetRootHandler)
	router.POST("/message", routes.PostMessageHandler)
	router.POST("/sensor", routes.CreateSensorHandler)
	router.GET("/sensor", routes.FindAllSensorsHandler)
	router.GET("/sensor/data", routes.FindAllSensorDataHandler)

	fmt.Println("Server started on port", defaultPort)

	err := router.Run(fmt.Sprintf(":%d", defaultPort))

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func showWelcomeMsg() {
	fmt.Println(color.CyanString(strings.Repeat("=", 25)))
	fmt.Println(color.GreenString(utils.CenterFormat("IoT Connect Broker", 25)))
	fmt.Println(color.CyanString(strings.Repeat("=", 25)))
}
