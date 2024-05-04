package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	checkSensor()
}

func TestGetSensorDataRoute(t *testing.T) {
	server := setupUdpServer() // inicia o servidor UDP para receber mensagens
	go server.Listen()         // inicia o servidor
	defer server.Close()

	cleanUp()
	router := setupRouter()
	w := httptest.NewRecorder()

	// adiciona sensor de teste
	body, _ := json.Marshal(gin.H{
		"address": "localhost:3399",
		"id":      "test_sensor",
	})
	req, _ := http.NewRequest("POST", "/sensor", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	// liga o sensor
	body, _ = json.Marshal(gin.H{
		"sensor_id": "test_sensor",
		"command":   "turn_on",
	})
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/message", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	// espera os dados do sensor
	time.Sleep(250 * time.Millisecond)

	// pede dados do sensor
	req, _ = http.NewRequest("GET", "/sensor/data", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	actual := []map[string]any{}
	json.Unmarshal(w.Body.Bytes(), &actual)

	assert.GreaterOrEqual(t, len(actual), 1)
	sensor := actual[0]

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "test_sensor", sensor["sensor_id"])
	assert.Equal(t, "temperature", sensor["name"])
	assert.NotNil(t, sensor["content"])
	assert.NotNil(t, sensor["created_at"])
	assert.NotNil(t, sensor["updated_at"])
}
