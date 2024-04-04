package main

import (
	"broker/routes"
	"broker/udp"
	"broker/utils"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
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
		fmt.Println("Received message:", msg)

		reply("Message received")
	})

	err := udpServer.Listen()

	if err != nil {
		log.Fatal("Error starting UDP server:", err)
	}
}

func handleServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", routes.GetRootHandler).Methods("GET")
	r.HandleFunc("/message", routes.PostMessageHandler).Methods("POST")
	r.HandleFunc("/sensor", routes.CreateSensorHandler).Methods("POST")
	r.HandleFunc("/sensor", routes.FindAllSensorsHandler).Methods("GET")

	fmt.Println("Server started on port", defaultPort)

	err := http.ListenAndServe(fmt.Sprintf(":%d", defaultPort), r)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func showWelcomeMsg() {
	fmt.Println(color.CyanString(strings.Repeat("=", 25)))
	fmt.Println(color.GreenString(utils.CenterFormat("IoT Connect Broker", 25)))
	fmt.Println(color.CyanString(strings.Repeat("=", 25)))
}
