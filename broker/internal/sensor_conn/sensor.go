package sensor_conn

import (
	"broker/internal/cmd"
	"errors"
	"net"
	"time"
)

const (
	timeout = 1 * time.Second

	handshakeSent     = "hello, sensor!"
	handshakeReceived = "hello, server!"
)

type Connection struct {
	Conn net.Conn
}

func New(addr string) (*Connection, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, err
	}

	if !ValidateSensorConnection(conn) {
		return nil, errors.New("validation failed")
	}

	return &Connection{Conn: conn}, nil
}

func Validate(addr string) error {
	conn, err := New(addr)
	if err != nil {
		return err
	}
	defer conn.Conn.Close()

	return nil
}

func (c *Connection) Send(content string) (string, error) {
	_, err := c.Conn.Write([]byte(content))
	if err != nil {
		return "", err
	}

	buff := make([]byte, 1024)
	n, err := c.Conn.Read(buff)
	if err != nil {
		return "", err
	}

	return string(buff[:n]), nil
}

func ValidateSensorConnection(conn net.Conn) bool {
	conn.Write([]byte(handshakeSent))
	buffer := make([]byte, len(handshakeReceived))
	n, err := conn.Read(buffer)

	if err != nil {
		return false
	}

	return string(buffer[:n]) == handshakeReceived
}

func Request(addr string, command *cmd.Cmd) (string, error) {
	conn, err := New(addr)
	if err != nil {
		return "", err
	}

	defer conn.Conn.Close()

	content, err := conn.Send(command.Decode())
	if err != nil {
		return "", err
	}

	return content, nil
}
