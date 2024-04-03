package storage

import (
	"broker/sensor"
	"sync"
)

type SensorStorage struct {
	sensors map[string]sensor.Sensor
	mutex   sync.RWMutex
}

var sensorStorage *SensorStorage

func GetSensorStorage() *SensorStorage {
	if sensorStorage == nil {
		sensorStorage = &SensorStorage{}
		sensorStorage.sensors = make(map[string]sensor.Sensor)
	}

	return sensorStorage
}

func (s *SensorStorage) AddSensor(sensor sensor.Sensor) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.sensors[sensor.Address] = sensor
}

func (s *SensorStorage) GetSensors() []sensor.Sensor {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	sensors := make([]sensor.Sensor, 0, len(s.sensors))
	for _, sensor := range s.sensors {
		sensors = append(sensors, sensor)
	}

	return sensors
}

func (s *SensorStorage) FindSensorByName(name string) *sensor.Sensor {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, sensor := range s.sensors {
		if sensor.Name == name {
			return &sensor
		}
	}

	return nil
}

func (s *SensorStorage) DeleteSensorByAddress(addr string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.sensors, addr)
}

func (s *SensorStorage) FindSensorByAddress(addr string) *sensor.Sensor {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if sensor, ok := s.sensors[addr]; ok {
		return &sensor
	}

	return nil
}
