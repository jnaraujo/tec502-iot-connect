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

func (s *SensorStorage) GetSensors() []map[string]string {
	var sensors []map[string]string = []map[string]string{}

	for addr, id := range s.addrs {
		sensors = append(sensors, map[string]string{
			"id":      id,
			"address": addr,
		})
	}

	return sensors
}

func (s *SensorStorage) FindSensorAddrById(id string) string {
	for addr, sensorId := range s.addrs {
		if sensorId == id {
			return addr
		}
	}

	return ""
}

func (s *SensorStorage) DeleteSensorByAddress(addr string) {
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
