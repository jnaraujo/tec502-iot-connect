package storage

import (
	"broker/internal/storage/responses"
	"broker/internal/storage/sensors"
)

func ClearAll() {
	responses.DeleteAll()
	sensors.DeleteAll()
}
