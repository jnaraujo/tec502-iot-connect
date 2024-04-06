package sensor_conn

import (
	"broker/internal/cmd_parser"
	"broker/internal/errors"
	"net"
	"os"
	"time"
)

const (
	timeout = 1 * time.Second

	handshakeSent     = "hello, sensor!"
	handshakeReceived = "hello, server!"
)

type SensorConn struct {
	Conn net.Conn
}

func NewSensorConn(addr string) (*SensorConn, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		switch {
		case os.IsTimeout(err):
			return nil, errors.ErrTimeout
		default:
			return nil, err
		}
	}

	if !ValidateSensorConnection(conn) {
		return nil, errors.ErrValidationFailed
	}

	return &SensorConn{
		Conn: conn,
	}, nil
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

/*
Cria um request para o sensor enviando o comando passado.

Exemplo de uso:

	conn, err := sensor.NewSensorConn("localhost:3333")
	if err != nil {
		log.Fatal(err)
	}

	response, err := conn.Request(cmd_parser.Cmd{
		ID:      "1",
		Command: "get",
		Content: "temperature",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
*/
func Request(addr string, cmd cmd_parser.Cmd) (string, error) {
	conn, err := NewSensorConn(addr)
	if err != nil {
		return "", err
	}

	defer conn.Conn.Close()

	_, err = conn.Conn.Write([]byte(cmd_parser.EncodeCmd(cmd)))
	if err != nil {
		return "", err
	}

	buff := make([]byte, 1024)
	n, err := conn.Conn.Read(buff)
	if err != nil {
		return "", err
	}

	return string(buff[:n]), nil
}
