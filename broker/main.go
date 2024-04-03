package main

import (
	"broker/routes"
	"broker/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

const (
	defaultPort = 8080
)

func main() {
	showWelcomeMsg()

	handleServer()
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
