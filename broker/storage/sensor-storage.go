package storage

import (
	"broker/sensor"
	"sync"
)

type SensorStorage struct {
	sensors []sensor.Sensor
	mutex   sync.RWMutex
}

var sensorStorage *SensorStorage

func GetSensorStorage() *SensorStorage {
	if sensorStorage == nil {
		sensorStorage = &SensorStorage{}
		sensorStorage.sensors = []sensor.Sensor{}
	}

	return sensorStorage
}

func (s *SensorStorage) AddSensor(sensor sensor.Sensor) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.sensors = append(s.sensors, sensor)
}

func (s *SensorStorage) GetSensors() []sensor.Sensor {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.sensors
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

	for i, sensor := range s.sensors {
		if sensor.Address == addr {
			s.sensors = append(s.sensors[:i], s.sensors[i+1:]...)
			break
		}
	}
}

func (s *SensorStorage) FindSensorByAddress(addr string) *sensor.Sensor {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, sensor := range s.sensors {
		if sensor.Address == addr {
			return &sensor
		}
	}

	return nil
}
