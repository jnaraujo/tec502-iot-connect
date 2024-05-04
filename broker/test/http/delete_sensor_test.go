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

func TestDeleteSensorRoute(t *testing.T) {
	cleanUp()
	router := setupRouter()
	w := httptest.NewRecorder()

	// cria um sensor de exemplo para ser deletado
	body, _ := json.Marshal(gin.H{
		"address": "localhost:3399",
		"id":      "test_sensor",
	})
	req, _ := http.NewRequest("POST", "/sensor", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	// deleta o sensor criado
	req, _ = http.NewRequest("DELETE", "/sensor/test_sensor", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	expected, _ := json.Marshal(gin.H{
		"message": "Sensor deletado.",
	})

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

func TestDeleteSensorRouteWithSensorNotFound(t *testing.T) {
	cleanUp()
	router := setupRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/sensor/test_sensor", nil)
	router.ServeHTTP(w, req)

	expected, _ := json.Marshal(gin.H{
		"message": "Sensor não encontrado.",
	})

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
