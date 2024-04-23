package udp_server

import (
	"errors"
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
	if conn == nil {
		return errors.New("connection is nil")
	}

	u.Conn = conn

	for {
		buffer := make([]byte, 1024)

		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			return err
		}
		if addr == nil {
			return errors.New("addr is nil")
		}

		go u.handler(addr.String(), string(buffer[:n]))
	}
}

func (u *UDPServer) HandleRequest(handler UDPServerHandler) {
	u.handler = handler
}

func (u *UDPServer) Close() {
	if u.Conn == nil {
		return
	}
	u.Conn.Close()
}
