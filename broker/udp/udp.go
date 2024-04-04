package udp

import (
	"fmt"
	"net"
)

type UDPServerHandler func(string, func(string) error)

type UDPServer struct {
	Conn net.PacketConn
	Addr string

	handler UDPServerHandler
}

func NewUDPServer(addr string) *UDPServer {
	return &UDPServer{
		Addr: addr,
	}
}

func (u *UDPServer) Listen() error {
	conn, err := net.ListenPacket("udp", u.Addr)

	if err != nil {
		return err
	}

	u.Conn = conn

	buffer := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFrom(buffer)

		fmt.Println("Received message from", addr, ":", string(buffer[:n]))

		if err != nil {
			return err
		}

		reply := u.makeReplayFunc(conn, addr)
		go u.handler(string(buffer[:n]), reply)
	}
}

func (u *UDPServer) makeReplayFunc(conn net.PacketConn, addr net.Addr) func(string) error {
	return func(msg string) error {
		_, err := conn.WriteTo([]byte(msg), addr)

		if err != nil {
			return err
		}

		return nil
	}
}

func (u *UDPServer) HandleRequest(handler UDPServerHandler) {
	u.handler = handler
}

func (u *UDPServer) Close() {
	u.Conn.Close()
}
