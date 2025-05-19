package web

import (
	"github.com/cifong/golang-toolsbox/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 首頁
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to golang-toolsbox!")
	})

	// 系統資訊 API
	r.GET("/api/system/info", handler.GetSystemInfo)

	return r
}
