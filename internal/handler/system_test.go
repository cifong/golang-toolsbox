package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	systemRoutes := router.Group("/api/v1/system")
	{
		systemRoutes.GET("/info", GetSystemInfo)
		systemRoutes.POST("/shutdown", ShutdownSystem)
	}
	return router
}

func TestGetSystemInfo(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/system/info", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), `"os":`)
	assert.Contains(t, rec.Body.String(), `"arch":`)
	assert.Contains(t, rec.Body.String(), `"version":`)
}

func TestShutdownSystem(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest(http.MethodPost, "/api/v1/system/shutdown", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	// 解析 JSON 回傳
	type Response struct {
		Message string `json:"message"`
	}

	var resp Response
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "System is shutting down...", resp.Message)
}
