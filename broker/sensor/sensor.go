package sensor

import (
	"broker/errors"
	"broker/types"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

type Sensor struct {
	Name    string
	Address string
	Conn    net.Conn
}

const (
	timeout = 1 * time.Second

	handshakeSent     = "hello, sensor!"
	handshakeReceived = "hello, server!"
)

func NewSensorConn(newSensor types.NewSensor) (*Sensor, error) {
	conn, err := net.DialTimeout("tcp", newSensor.Address, timeout)

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
		Name:    newSensor.Name,
		Address: newSensor.Address,
		Conn:    conn,
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

func (s *Sensor) OnDataReceived(
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

func (s *Sensor) Close() error {
	return s.Conn.Close()
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

func (s *Sensor) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"name":    s.Name,
		"address": s.Address,
	})
}
