package storage

import (
	"broker/internal/time"
)

type Response struct {
	SensorID  string    `json:"sensor_id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SensorResponseStorage struct {
	data map[string]Response
}

var storage *SensorResponseStorage

func GetSensorResponseStorage() *SensorResponseStorage {
	if storage == nil {
		storage = &SensorResponseStorage{
			data: make(map[string]Response),
		}
	}

	return storage
}

func (s *SensorResponseStorage) Create(sensorID, name, content string) Response {
	response := Response{
		SensorID:  sensorID,
		Name:      name,
		Content:   content,
		CreatedAt: *time.NewTimeNow(),
	}

	s.data[sensorID] = response

	return response
}

func (s *SensorResponseStorage) FindAll() []Response {
	responses := []Response{}

	for _, resp := range s.data {
		responses = append(responses, resp)
	}

	return responses
}

func (s *SensorResponseStorage) FindBySensorId(sensorId string) Response {
	return s.data[sensorId]
}

func (s *SensorResponseStorage) DeleteBySensorId(sensorId string) {
	delete(s.data, sensorId)
}

func (s *SensorResponseStorage) UpdateContent(sensorId, content string) {
	response := s.data[sensorId]

	response.Content = content
	response.UpdatedAt = *time.NewTimeNow()

	s.data[sensorId] = response
}
