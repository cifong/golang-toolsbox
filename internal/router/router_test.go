package router_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/cifong/golang-toolsbox/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestMain(m *testing.M) {
	dir, _ := os.Getwd()
	for !fileExists(filepath.Join(dir, "go.mod")) && dir != filepath.Dir(dir) {
		dir = filepath.Dir(dir)
	}
	_ = os.Chdir(dir)
	os.Exit(m.Run())
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func TestPageRoutes(t *testing.T) {
	r := router.SetupRouter()

	tests := []struct {
		path         string
		expectedCode int
		expectedBody string
	}{
		{"/", http.StatusOK, "<!DOCTYPE html"},     // 檢查首頁
		{"/todo", http.StatusOK, "<!DOCTYPE html"}, // 檢查 ToDo 頁
	}

	for _, tt := range tests {
		req, _ := http.NewRequest("GET", tt.path, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, tt.expectedCode, w.Code)
		assert.Contains(t, w.Body.String(), tt.expectedBody)
	}
}

func TestSystemInfoAPI(t *testing.T) {
	r := router.SetupRouter()

	req, _ := http.NewRequest("GET", "/api/v1/system/info", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "os")
	assert.Contains(t, w.Body.String(), "arch")
	assert.Contains(t, w.Body.String(), "version")
	assert.Contains(t, w.Body.String(), "cpu_usage")
	assert.Contains(t, w.Body.String(), "total_memory")
	assert.Contains(t, w.Body.String(), "used_memory")
}

func TestShutdownAPI(t *testing.T) {
	r := router.SetupRouter()

	req, _ := http.NewRequest("POST", "/api/v1/system/shutdown", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "System is shutting down")
}
