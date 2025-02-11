package tests

import (
	"commonplace-go/internal/api"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHelloHandler(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/hello", nil)

	// Execute
	api.HelloHandler(c)

	// Assert
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}

	if w.Body.String() != "Hello, World!" {
		t.Errorf("expected body %s, got %s", "Hello, World!", w.Body.String())
	}
}

func TestHealthCheckHandler(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/health", nil)

	// Execute
	api.HealthCheckHandler(c)

	// Assert
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}
	if w.Body.String() != "OK" {
		t.Errorf("expected body %s, got %s", "OK", w.Body.String())
	}
}
