package udp_server

import (
	"net"
)

// UDPServerHandler é uma função que lida com as mensagens recebidas pelo servidor UDP.
type UDPServerHandler func(addr, content string)

type UDPServer struct {
	Conn net.PacketConn // Conn é a conexão do servidor UDP.
	Addr string         // Addr é o endereço do servidor UDP.

	handler UDPServerHandler // handler é a função que lida com as mensagens recebidas pelo servidor UDP.
}

// NewUDPServer cria uma nova instância de UDPServer.
func NewUDPServer(addr string) *UDPServer {
	return &UDPServer{
		Addr: addr,
	}
}

// Listen inicia o servidor UDP.
func (u *UDPServer) Listen() error {
	conn, err := net.ListenPacket("udp", u.Addr)
	if err != nil {
		return err
	}
	u.Conn = conn

	for {
		buffer := make([]byte, 1024)

		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			return err
		}

		go u.handler(addr.String(), string(buffer[:n]))
	}
}

// HandleRequest define a função que lida com as mensagens recebidas pelo servidor UDP.
func (u *UDPServer) HandleRequest(handler UDPServerHandler) {
	u.handler = handler
}

// Close fecha a conexão do servidor UDP.
func (u *UDPServer) Close() {
	u.Conn.Close()
}
