package main

import (
	"broker/internal/http"
	"broker/internal/udp_server"
	"broker/utils"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func main() {
	fmt.Println(color.CyanString(strings.Repeat("=", 25)))
	fmt.Println(color.GreenString(utils.CenterFormat("IoT Connect Broker", 25)))
	fmt.Println(color.CyanString(strings.Repeat("=", 25)))

	go udp_server.NewServer("", 5310)
	http.NewServer("", 8080)
}
