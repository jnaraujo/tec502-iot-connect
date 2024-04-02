package main

import (
	"broker/sensor"
	"broker/utils"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

func main() {
	showWelcomeMsg()

	conn, err := sensor.NewSensorConn("localhost:3333")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	defer conn.Close()

	conn.Send("Hello, server!")

	counter := 0
	conn.OnDataReceived(func(data string) {
		fmt.Println("Received:", data)
		time.Sleep(1 * time.Second)
		counter++
		conn.Send(fmt.Sprintf("Hello, server! %d", counter))
	}, 1024)

	// server.Init()
}

func showWelcomeMsg() {
	fmt.Println(color.CyanString(strings.Repeat("=", 25)))
	fmt.Println(color.GreenString(utils.CenterFormat("IoT Connect Broker", 25)))
	fmt.Println(color.CyanString(strings.Repeat("=", 25)))
}
