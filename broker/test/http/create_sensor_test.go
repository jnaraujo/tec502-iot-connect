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

func TestCreateSensorRoute(t *testing.T) {
	cleanUp()
	router := setupRouter()
	w := httptest.NewRecorder()

	body, _ := json.Marshal(gin.H{
		"address": "localhost:3399",
		"id":      "test_sensor",
	})
	req, _ := http.NewRequest("POST", "/sensor", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	expected, _ := json.Marshal(gin.H{
		"message": "Sensor criado",
	})

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

func TestCreateSensorRouteWithOfflineSensor(t *testing.T) {
	cleanUp()
	router := setupRouter()
	w := httptest.NewRecorder()

	body, _ := json.Marshal(gin.H{
		"address": "localhost:9999",
		"id":      "test_sensor",
	})
	req, _ := http.NewRequest("POST", "/sensor", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	expected, _ := json.Marshal(gin.H{
		"message": "Erro ao se conectar com o sensor",
	})

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

func TestCreateSensorRouteWithMissingAddress(t *testing.T) {
	cleanUp()
	router := setupRouter()
	w := httptest.NewRecorder()

	body, _ := json.Marshal(gin.H{
		"id": "test_sensor",
	})
	req, _ := http.NewRequest("POST", "/sensor", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	expected, _ := json.Marshal(gin.H{
		"message": "Corpo da requisição é inválido",
	})

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

func TestCreateSensorRouteWithMissingId(t *testing.T) {
	cleanUp()
	router := setupRouter()
	w := httptest.NewRecorder()

	body, _ := json.Marshal(gin.H{
		"address": "localhost:3399",
	})
	req, _ := http.NewRequest("POST", "/sensor", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	expected, _ := json.Marshal(gin.H{
		"message": "Corpo da requisição é inválido",
	})

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
