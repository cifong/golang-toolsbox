package router

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
	systemRoutes := r.Group("/api/v1/system")
	{
		systemRoutes.GET("/info", handler.GetSystemInfo)       // 取得系統資訊
		systemRoutes.POST("/shutdown", handler.ShutdownSystem) // 關機
	}

	r.Static("/todo", "./web/todo")

	return r
}
