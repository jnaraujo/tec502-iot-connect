package storage

import (
	"broker/sensor"
	"sync"
)

var mutex = sync.RWMutex{}

type SensorStorage struct {
	sensors []sensor.Sensor
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
	mutex.Lock()
	defer mutex.Unlock()
	s.sensors = append(s.sensors, sensor)
}

func (s *SensorStorage) GetSensors() []sensor.Sensor {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.sensors
}

func (s *SensorStorage) FindSensorByAddress(addr string) *sensor.Sensor {
	mutex.RLock()
	defer mutex.RUnlock()

	for _, sensor := range s.sensors {
		if sensor.Address == addr {
			return &sensor
		}
	}

	return nil
}
