// Este pacote é responsável por estabelecer a conexão com o sensor e enviar comandos para ele.
package sensor_conn

import (
	"broker/internal/cmd"
	"errors"
	"net"
	"time"
)

const (
	timeout = 1 * time.Second // tempo limite para a conexão com o sensor

	handshakeSent     = "hello, sensor!" // handshake enviado para o sensor
	handshakeReceived = "hello, server!" // handshake recebido do sensor
)

type Connection struct {
	Conn net.Conn
}

// New cria uma nova conexão com o sensor.
func New(addr string) (*Connection, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, err
	}

	if !validateSensorConnection(conn) {
		return nil, errors.New("validation failed")
	}

	return &Connection{Conn: conn}, nil
}

// ValidateSensorConnection verifica se a conexão com o sensor é válida.
// Para isso, é enviado um handshake para o sensor e é esperado que o sensor responda com um handshake também.
func validateSensorConnection(conn net.Conn) bool {
	conn.Write([]byte(handshakeSent))
	buffer := make([]byte, len(handshakeReceived))
	n, err := conn.Read(buffer)

	if err != nil {
		return false
	}

	return string(buffer[:n]) == handshakeReceived
}

// Validate verifica se a conexão com o sensor é válida.
// Ou seja, se o sensor está respondendo.
func Validate(addr string) error {
	conn, err := New(addr)
	if err != nil {
		return err
	}
	defer conn.Conn.Close()

	return nil
}

// Send envia uma mensagem para o sensor e retorna a resposta.
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

// Request envia um comando para o sensor e retorna a resposta.
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
