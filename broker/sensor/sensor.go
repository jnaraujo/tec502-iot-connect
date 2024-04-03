package sensor

import (
	"encoding/json"
	"errors"
	"net"
)

type Sensor struct {
	Name    string
	Address string
	Conn    net.Conn
}

type NewSensor struct {
	Name    string
	Address string
}

func NewSensorConn(newSensor NewSensor) (*Sensor, error) {
	conn, err := net.Dial("tcp", newSensor.Address)
	if err != nil {
		return nil, err
	}

	if !ValidateHandshake(conn) {
		return nil, errors.New("handshake failed")
	}

	return &Sensor{
		Name:    newSensor.Name,
		Address: newSensor.Address,
		Conn:    conn,
	}, nil
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

func (s *Sensor) Send(data string) (int, error) {
	return s.Conn.Write([]byte(data))
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
