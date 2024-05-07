// Este pacote é responsável por armazenar os sensores..
package sensors

import "sync"

type SensorStorage struct {
	mu    sync.RWMutex
	addrs map[string]string
}

// storage é uma instância de SensorStorage que armazena os sensores.
var storage *SensorStorage = &SensorStorage{
	addrs: map[string]string{},
	mu:    sync.RWMutex{},
}

// AddSensor adiciona um sensor ao armazenamento.
func AddSensor(id string, addr string) {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	storage.addrs[addr] = id
}

// Sensor é uma estrutura que representa um sensor.
type Sensor struct {
	Id      string `json:"id"`
	Address string `json:"address"`
}

// FindSensors retorna todos os sensores.
func FindSensors() []Sensor {
	storage.mu.RLock()
	defer storage.mu.RUnlock()

	var sensors []Sensor = []Sensor{}

	for addr, id := range storage.addrs {
		sensors = append(sensors, Sensor{
			Id:      id,
			Address: addr,
		})
	}

	return sensors
}

// DoesSensorExists verifica se um sensor existe.
func DoesSensorExists(id, addr string) bool {
	storage.mu.RLock()
	defer storage.mu.RUnlock()

	sensor_addr_exists := FindSensorAddrById(id) != ""
	sensor_id_exists := FindSensorIdByAddress(addr) != ""
	return sensor_addr_exists || sensor_id_exists
}

// FindSensorAddrById encontra o endereço de um sensor pelo seu ID.
func FindSensorAddrById(id string) string {
	storage.mu.RLock()
	defer storage.mu.RUnlock()

	for addr, sensorId := range storage.addrs {
		if sensorId == id {
			return addr
		}
	}

	return ""
}

// DeleteSensorBySensorId deleta um sensor pelo seu ID.
func DeleteSensorBySensorId(sensorId string) {
	addr := FindSensorAddrById(sensorId)

	storage.mu.Lock()
	delete(storage.addrs, addr)
	storage.mu.Unlock()
}

// FindSensorIdByAddress encontra o ID de um sensor pelo seu endereço.
func FindSensorIdByAddress(addr string) string {
	storage.mu.RLock()
	defer storage.mu.RUnlock()

	for sensorAddr, sensorId := range storage.addrs {
		if sensorAddr == addr {
			return sensorId
		}
	}

	return ""
}

// Deleta todos os Sensores
func DeleteAll() {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	for k := range storage.addrs {
		delete(storage.addrs, k)
	}
}
