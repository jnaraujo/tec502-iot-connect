package udp_server

import (
	"net"
	"testing"
)

func TestUdpServer(t *testing.T) {
	udpServer := NewUDPServer(":3333")
	defer udpServer.Close()

	udpServer.HandleRequest(func(msg string, reply func(string) error) {
		reply("Message received")
	})
	go udpServer.Listen()

	udpClient, err := udpClient()
	if err != nil {
		t.Error("Error creating UDP client:", err)
	}

	defer udpClient.Close()

	_, err = udpClient.Write([]byte("Hello, world!"))

	if err != nil {
		t.Error("Error sending message to UDP server:", err)
	}

	buffer := make([]byte, 1024)
	n, err := udpClient.Read(buffer)

	if err != nil {
		t.Error("Error reading response from UDP server:", err)
	}

	if string(buffer[:n]) != "Message received" {
		t.Error("Unexpected response from UDP server:", string(buffer[:n]))
	}
}

func udpClient() (net.Conn, error) {
	udpClient, err := net.Dial("udp", "localhost:3333")
	return udpClient, err
}
