package sensors

type SensorStorage struct {
	addrs map[string]string
}

var storage *SensorStorage = &SensorStorage{
	map[string]string{},
}

func AddSensor(id string, addr string) {
	storage.addrs[addr] = id
}

type Sensor struct {
	Id      string `json:"id"`
	Address string `json:"address"`
}

func FindSensors() []Sensor {
	var sensors []Sensor = []Sensor{}

	for addr, id := range storage.addrs {
		sensors = append(sensors, Sensor{
			Id:      id,
			Address: addr,
		})
	}

	return sensors
}

func DoesSensorExists(id, addr string) bool {
	sensor := FindSensorAddrById(id)
	if sensor != "" {
		return true
	}

	sensor = FindSensorIdByAddress(id)
	return sensor != ""
}

func FindSensorAddrById(id string) string {
	for addr, sensorId := range storage.addrs {
		if sensorId == id {
			return addr
		}
	}

	return ""
}

func DeleteSensorBySensorId(sensorId string) {
	addr := FindSensorAddrById(sensorId)
	delete(storage.addrs, addr)
}

func FindSensorIdByAddress(addr string) string {
	for sensorAddr, sensorId := range storage.addrs {
		if sensorAddr == addr {
			return sensorId
		}
	}

	return ""
}
