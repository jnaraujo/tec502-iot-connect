package storage

type SensorStorage struct {
	addrs map[string]string
}

var sensorStorage *SensorStorage

func GetSensorStorage() *SensorStorage {
	if sensorStorage == nil {
		sensorStorage = &SensorStorage{
			addrs: make(map[string]string),
		}
	}

	return sensorStorage
}

func (s *SensorStorage) AddSensor(id string, addr string) {
	s.addrs[addr] = id
}

type Sensor struct {
	Id      string `json:"id"`
	Address string `json:"address"`
}

func (s *SensorStorage) FindSensors() []Sensor {
	var sensors []Sensor = []Sensor{}

	for addr, id := range s.addrs {
		sensors = append(sensors, Sensor{
			Id:      id,
			Address: addr,
		})
	}

	return sensors
}

func (s *SensorStorage) DoesSensorExists(id, addr string) bool {
	sensor := s.FindSensorAddrById(id)
	if sensor != "" {
		return true
	}

	sensor = s.FindSensorIdByAddress(id)
	return sensor != ""
}

func (s *SensorStorage) FindSensorAddrById(id string) string {
	for addr, sensorId := range s.addrs {
		if sensorId == id {
			return addr
		}
	}

	return ""
}

func (s *SensorStorage) DeleteSensorBySensorId(sensorId string) {
	addr := s.FindSensorAddrById(sensorId)
	delete(s.addrs, addr)
}

func (s *SensorStorage) FindSensorIdByAddress(addr string) string {
	for sensorAddr, sensorId := range s.addrs {
		if sensorAddr == addr {
			return sensorId
		}
	}

	return ""
}
