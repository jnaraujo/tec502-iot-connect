package server

import (
	"broker/server/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	DefaultPort = 8080
)

func Init() {
	r := mux.NewRouter()

	r.HandleFunc("/", routes.GetRootHandler).Methods("GET")
	r.HandleFunc("/message", routes.PostMessageHandler).Methods("POST")
	r.HandleFunc("/sensor", routes.CreateSensorHandler).Methods("POST")
	r.HandleFunc("/sensor", routes.FindAllSensorsHandler).Methods("GET")

	fmt.Println("Server started on port", DefaultPort)

	err := http.ListenAndServe(fmt.Sprintf(":%d", DefaultPort), r)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
