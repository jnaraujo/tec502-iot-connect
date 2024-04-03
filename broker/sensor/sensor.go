package sensor

import (
	"broker/errors"
	"fmt"
	"net"
	"os"
	"time"
)

type Sensor struct {
	Conn net.Conn
}

const (
	timeout = 1 * time.Second

	handshakeSent     = "hello, sensor!"
	handshakeReceived = "hello, server!"
)

func NewSensorConn(addr string) (*Sensor, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)

	if err != nil {
		switch {
		case os.IsTimeout(err):
			return nil, errors.ErrTimeout
		default:
			return nil, err
		}
	}

	if !ValidateConnection(conn) {
		return nil, errors.ErrValidationFailed
	}

	return &Sensor{
		Conn: conn,
	}, nil
}

func ValidateConnection(conn net.Conn) bool {
	conn.Write([]byte(handshakeSent))
	buffer := make([]byte, len(handshakeReceived))
	n, err := conn.Read(buffer)

	if err != nil {
		return false
	}

	return string(buffer[:n]) == handshakeReceived
}

func (s *Sensor) Request(command string, content string) (string, error) {
	err := s.Send(command, content)
	if err != nil {
		return "nil", err
	}

	buffer := make([]byte, 1024)
	n, err := s.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer[:n]), nil
}

func (s *Sensor) Send(command string, content string) error {
	_, err := s.Conn.Write([]byte(fmt.Sprintf(
		"Cmd: %s\n\n"+
			"%s",
		command, content,
	)))
	return err
}

func (s *Sensor) Read(data []byte) (int, error) {
	return s.Conn.Read(data)
}

func (s *Sensor) Close() {
	s.Conn.Close()
}
