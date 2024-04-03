package storage

import (
	"sync"
)

type SensorStorage struct {
	addrs map[string]string
	mutex sync.RWMutex
}

var sensorStorage *SensorStorage

func GetSensorStorage() *SensorStorage {
	if sensorStorage == nil {
		sensorStorage = &SensorStorage{}
		sensorStorage.addrs = make(map[string]string)
	}

	return sensorStorage
}

func (s *SensorStorage) AddSensor(name string, addr string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.addrs[addr] = name
}

func (s *SensorStorage) GetSensors() []map[string]string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var sensors []map[string]string

	for addr, name := range s.addrs {
		sensors = append(sensors, map[string]string{
			"name":    name,
			"address": addr,
		})
	}

	return sensors
}

func (s *SensorStorage) FindSensorAddrByName(name string) string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for addr, sensorName := range s.addrs {
		if sensorName == name {
			return addr
		}
	}

	return ""
}

func (s *SensorStorage) DeleteSensorByAddress(addr string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.addrs, addr)
}

func (s *SensorStorage) FindSensorNameByAddress(addr string) string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for sensorAddr, sensorName := range s.addrs {
		if sensorAddr == addr {
			return sensorName
		}
	}

	return ""
}
