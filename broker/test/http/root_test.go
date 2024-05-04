package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRootRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	expected, _ := json.Marshal(gin.H{
		"message": "Welcome to the Broker API",
	})

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
