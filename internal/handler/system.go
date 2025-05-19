package handler

import (
	"net/http"

	"github.com/cifong/golang-toolsbox/internal/system"
	"github.com/gin-gonic/gin"
)

func GetSystemInfo(c *gin.Context) {
	info, err := system.GetSystemInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, info)
}

func ShutdownSystem(c *gin.Context) {
	// 回應 JSON 後再背景執行關機
	go func() {
		_ = system.Shutdown()
	}()

	c.JSON(http.StatusOK, gin.H{"message": "System is shutting down..."})
}
