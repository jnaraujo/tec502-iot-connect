package main

import (
	"broker/internal/http"
	"broker/internal/udpserver"
	"broker/utils"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func main() {
	fmt.Println(color.CyanString(strings.Repeat("=", 25)))
	fmt.Println(color.GreenString(utils.CenterFormat("IoT Connect Broker", 25)))
	fmt.Println(color.CyanString(strings.Repeat("=", 25)))

	go udpserver.NewServer("0.0.0.0", 5310)
	http.NewServer("0.0.0.0", 8080)
}
