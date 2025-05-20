package router

import (
	"github.com/cifong/golang-toolsbox/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 載入模板與靜態資源
	router.LoadHTMLGlob("web/templates/*")
	router.Static("/css", "./web/static/css")
	router.Static("/js", "./web/static/js")

	// 頁面路由
	pageRoutes := router.Group("/")
	{
		pageRoutes.GET("/", handler.RenderIndex)
		pageRoutes.GET("/todo", handler.RenderTodo)
	}

	// API 路由
	api := router.Group("/api/v1")
	{
		systemRoutes := api.Group("/system")
		{
			systemRoutes.GET("/info", handler.GetSystemInfo)
			systemRoutes.POST("/shutdown", handler.ShutdownSystem)
		}
	}

	return router
}
