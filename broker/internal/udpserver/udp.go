package udpserver

import (
	"net"
)

type UDPServerHandler func(addr, content string)

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

		if err != nil {
			return err
		}

		go u.handler(addr.String(), string(buffer[:n]))
	}
}

func (u *UDPServer) HandleRequest(handler UDPServerHandler) {
	u.handler = handler
}

func (u *UDPServer) Close() {
	u.Conn.Close()
}
