package router

import (
	"github.com/cifong/golang-toolsbox/internal/handler/api"
	"github.com/cifong/golang-toolsbox/internal/handler/web"
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
		pageRoutes.GET("/", web.RenderIndex)
		pageRoutes.GET("/todo", web.RenderTodo)
	}

	// API 路由
	apiRoutes := router.Group("/api/v1")
	{
		systemRoutes := apiRoutes.Group("/system")
		{
			systemRoutes.GET("/info", api.GetSystemInfo)
			systemRoutes.POST("/shutdown", api.ShutdownSystem)
		}
	}

	return router
}
