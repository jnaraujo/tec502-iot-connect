package responses

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

var storage *SensorResponseStorage = &SensorResponseStorage{
	data: make(map[string]Response),
}

func Create(sensorID, name, content string) Response {
	response := Response{
		SensorID:  sensorID,
		Name:      name,
		Content:   content,
		CreatedAt: *time.NewTimeNow(),
		UpdatedAt: *time.NewTimeNow(),
	}

	storage.data[sensorID] = response

	return response
}

func FindAll() []Response {
	responses := []Response{}

	for _, resp := range storage.data {
		responses = append(responses, resp)
	}

	return responses
}

func FindBySensorId(sensorId string) Response {
	return storage.data[sensorId]
}

func DeleteBySensorId(sensorId string) {
	delete(storage.data, sensorId)
}

func UpdateContent(sensorId, content string) {
	response := storage.data[sensorId]

	response.Content = content
	response.UpdatedAt = *time.NewTimeNow()

	storage.data[sensorId] = response
}
