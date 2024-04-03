package sensor

import (
	"errors"
	"net"
)

type SensorConn struct {
	Conn net.Conn
}

func NewSensorConn(address string) (*SensorConn, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	if !ValidateHandshake(conn) {
		return nil, errors.New("handshake failed")
	}

	return &SensorConn{Conn: conn}, nil
}

func ValidateHandshake(conn net.Conn) bool {
	conn.Write([]byte("hello, sensor!"))
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)

	if err != nil {
		return false
	}

	return string(buffer[:n]) == "hello, server!"
}

func (s *SensorConn) OnDataReceived(
	handle func(string), bufferSize_optional ...int,
) {
	bufferSize := 1024

	if len(bufferSize_optional) > 0 {
		bufferSize = bufferSize_optional[0]
	}

	for {
		buffer := make([]byte, bufferSize)
		n, err := s.Conn.Read(buffer)
		if err != nil {
			return
		}

		handle(string(buffer[:n]))
	}
}

func (s *SensorConn) Close() error {
	return s.Conn.Close()
}

func (s *SensorConn) Send(data string) (int, error) {
	return s.Conn.Write([]byte(data))
}

func (s *SensorConn) Read(data []byte) (int, error) {
	return s.Conn.Read(data)
}
