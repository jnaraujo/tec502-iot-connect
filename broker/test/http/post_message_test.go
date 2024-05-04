package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	checkSensor()
}

func TestPostMessageRoute(t *testing.T) {
	cleanUp()
	router := setupRouter()
	w := httptest.NewRecorder()

	// cria um sensor de exemplo
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

	expected, _ := json.Marshal(gin.H{
		"message": "Ar Condicionado foi ligado",
	})
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expected), w.Body.String())

	// alterar a temperatura
	body, _ = json.Marshal(gin.H{
		"sensor_id": "test_sensor",
		"command":   "set_temp",
		"content":   "50",
	})
	res, _ := http.NewRequest("POST", "/message", bytes.NewBuffer(body))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, res)

	expected, _ = json.Marshal(gin.H{
		"message": "Temperature set to 50.0",
	})
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

func TestPostMessageRouteSensorOffline(t *testing.T) {
	cleanUp()
	router := setupRouter()
	w := httptest.NewRecorder()

	// cria um sensor de exemplo
	body, _ := json.Marshal(gin.H{
		"address": "localhost:3399",
		"id":      "test_sensor",
	})
	req, _ := http.NewRequest("POST", "/sensor", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	// desligar o sensor
	body, _ = json.Marshal(gin.H{
		"sensor_id": "test_sensor",
		"command":   "turn_off",
	})
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/message", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	expected, _ := json.Marshal(gin.H{
		"message": "Ar Condicionado foi desligado",
	})
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expected), w.Body.String())

	// alterar a temperatura
	body, _ = json.Marshal(gin.H{
		"sensor_id": "test_sensor",
		"command":   "set_temp",
		"content":   "50",
	})
	res, _ := http.NewRequest("POST", "/message", bytes.NewBuffer(body))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, res)

	expected, _ = json.Marshal(gin.H{
		"message": "O sensor está desligado",
	})
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
