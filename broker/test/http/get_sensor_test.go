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

func TestGetSensors(t *testing.T) {
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

	req, _ = http.NewRequest("GET", "/sensor", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	actual := map[string][]map[string]any{}
	json.Unmarshal(w.Body.Bytes(), &actual)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 1, len(actual["sensors"]))
	assert.Equal(t, "test_sensor", actual["sensors"][0]["id"])
	assert.Equal(t, "localhost:3399", actual["sensors"][0]["address"])
	assert.NotNil(t, actual["sensors"][0]["is_online"])
}
