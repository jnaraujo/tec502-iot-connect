package main

import (
	"broker/internal/http"
	"broker/internal/udp_server"
	"broker/internal/utils"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

/*
A função main é o ponto de entrada de um programa Go. Ela é responsável por iniciar a execução do programa.

Neste caso, a função main inicia um servidor UDP e um servidor HTTP.
*/
func main() {
	fmt.Println(color.CyanString(strings.Repeat("=", 25)))
	fmt.Println(color.GreenString(utils.CenterFormat("IoT Connect Broker", 25)))
	fmt.Println(color.CyanString(strings.Repeat("=", 25)))

	server := udp_server.NewServer("0.0.0.0", 5310)
	defer server.Close()
	go server.Listen()

	http.NewServer("0.0.0.0", 8080)
}
