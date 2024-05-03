// Este pacote é responsável por armazenar as respostas dos sensores.
package responses

import (
	"broker/internal/queue"
	"broker/internal/time"
	"sync"
)

type Response struct {
	SensorID  string                `json:"sensor_id"`
	Name      string                `json:"name"`
	Content   *queue.Queue[float64] `json:"content"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

type SensorResponseStorage struct {
	mu   sync.RWMutex // mu é um mutex para controlar o acesso concorrente ao mapa data.
	data map[string]Response
}

// storage é uma instância de SensorResponseStorage que armazena as respostas dos sensores.
var storage *SensorResponseStorage = &SensorResponseStorage{
	data: make(map[string]Response),
}

// Create cria uma nova resposta para um sensor.
func Create(sensorID, name string) Response {
	storage.mu.Lock()
	defer storage.mu.Unlock()

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

// FindAll retorna todas as respostas dos sensores.
func FindAll() []Response {
	storage.mu.RLock()
	defer storage.mu.RUnlock()

	responses := []Response{}
	for _, resp := range storage.data {
		responses = append(responses, resp)
	}
	return responses
}

// FindBySensorId retorna a resposta de um sensor pelo seu ID.
func FindBySensorId(sensorId string) Response {
	storage.mu.RLock()
	defer storage.mu.RUnlock()

	return storage.data[sensorId]
}

// DeleteBySensorId deleta a resposta de um sensor pelo seu ID.
func DeleteBySensorId(sensorId string) {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	delete(storage.data, sensorId)
}

// AddContent adiciona um conteúdo à resposta de um sensor.
func AddContent(sensorId string, data float64) {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	response := storage.data[sensorId]
	response.Content.Add(data)
	response.UpdatedAt = *time.NewTimeNow()
	storage.data[sensorId] = response
}
