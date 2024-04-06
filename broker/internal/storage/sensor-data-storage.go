package storage

import (
	"broker/internal/time"
)

type SensorData struct {
	ID         int       `json:"id"`
	SensorID   string    `json:"sensor_id"`
	Command    string    `json:"command"`
	Content    string    `json:"content"`
	Response   string    `json:"response"`
	CreatedAt  time.Time `json:"created_at"`
	ReceivedAt time.Time `json:"received_at,omitempty"`
}

type SensorDataStorage struct {
	data []SensorData
}

var sensorDataStorage *SensorDataStorage

func GetSensorDataStorage() *SensorDataStorage {
	if sensorDataStorage == nil {
		sensorDataStorage = &SensorDataStorage{
			data: []SensorData{},
		}
	}

	return sensorDataStorage
}

func (s *SensorDataStorage) Create(sensorID, command string, content string) *SensorData {
	sensor := SensorData{
		ID:        len(s.data) + 1,
		SensorID:  sensorID,
		Command:   command,
		Content:   content,
		CreatedAt: *time.NewTimeNow(),
	}

	s.data = append(s.data, sensor)

	return &sensor
}

func (s *SensorDataStorage) FindAll() []SensorData {
	return s.data
}

func (s *SensorDataStorage) FindByID(id int) *SensorData {
	for _, data := range s.data {
		if data.ID == id {
			return &data
		}
	}

	return nil
}

func (s *SensorDataStorage) UpdateResponse(id int, response string) {
	for i, data := range s.data {
		if data.ID == id {
			currentData := s.data[i]
			currentData.Response = response
			currentData.ReceivedAt = *time.NewTimeNow()

			s.data[i] = currentData
			return
		}
	}
}
