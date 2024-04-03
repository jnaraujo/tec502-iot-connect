package main

import (
	"broker/server"
	"broker/utils"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func main() {
	showWelcomeMsg()
	server.Init()
}

func showWelcomeMsg() {
	fmt.Println(color.CyanString(strings.Repeat("=", 25)))
	fmt.Println(color.GreenString(utils.CenterFormat("IoT Connect Broker", 25)))
	fmt.Println(color.CyanString(strings.Repeat("=", 25)))
}
