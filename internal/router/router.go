package router

import (
	"github.com/cifong/golang-toolsbox/internal/handler/api"
	"github.com/cifong/golang-toolsbox/internal/handler/web"
	"github.com/cifong/golang-toolsbox/internal/handler/websocket"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 載入模板與靜態資源
	router.LoadHTMLGlob("web/templates/*")
	router.Static("/css", "./web/static/css")
	router.Static("/js", "./web/static/js")
	// icon 與 manifest 檔案
	router.StaticFile("/favicon.ico", "./web/static/favicon.ico")
	router.StaticFile("/favicon.svg", "./web/static/favicon.svg")
	router.StaticFile("/favicon-96x96.png", "./web/static/favicon-96x96.png")
	router.StaticFile("/apple-touch-icon.png", "./web/static/apple-touch-icon.png")
	router.StaticFile("/site.webmanifest", "./web/static/site.webmanifest")
	router.StaticFile("/web-app-manifest-192x192.png", "./web/static/web-app-manifest-192x192.png")
	router.StaticFile("/web-app-manifest-512x512.png", "./web/static/web-app-manifest-512x512.png")

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

	// WebSocket 路由
	wsRoutes := router.Group("/ws/v1")
	{
		systemRoutes := wsRoutes.Group("/system")
		{
			systemRoutes.GET("/info", websocket.GetSystemInfoWebSocket)
		}
	}
	return router
}
