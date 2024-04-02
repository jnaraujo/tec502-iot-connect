package sensor

import "net"

type SensorConn struct {
	Conn net.Conn
}

func NewSensorConn(address string) (*SensorConn, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return &SensorConn{Conn: conn}, nil
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
