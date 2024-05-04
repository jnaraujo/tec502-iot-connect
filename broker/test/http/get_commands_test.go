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

func TestGetCommandsRoute(t *testing.T) {
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

	req, _ = http.NewRequest("GET", "/sensor/commands/test_sensor", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	actual := map[string][]string{}
	json.Unmarshal(w.Body.Bytes(), &actual)

	expectedCommand := []string{"turn_on", "turn_off", "set_temp", "not_found", "set_heat", "set_cool"}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.ElementsMatch(t, expectedCommand, actual["commands"])
}
