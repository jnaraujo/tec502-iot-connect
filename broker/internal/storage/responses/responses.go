package responses

import (
	"broker/internal/queue"
	"broker/internal/time"
)

type Response struct {
	SensorID  string                `json:"sensor_id"`
	Name      string                `json:"name"`
	Content   *queue.Queue[float64] `json:"content"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

type SensorResponseStorage struct {
	data map[string]Response
}

var storage *SensorResponseStorage = &SensorResponseStorage{
	data: make(map[string]Response),
}

func Create(sensorID, name string) Response {
	response := Response{
		SensorID:  sensorID,
		Name:      name,
		Content:   queue.New[float64](20),
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

func AddContent(sensorId string, data float64) {
	response := storage.data[sensorId]

	response.Content.Add(data)
	response.UpdatedAt = *time.NewTimeNow()

	storage.data[sensorId] = response
}
