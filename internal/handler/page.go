package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RenderIndex 渲染首頁
func RenderIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// RenderTodo 渲染 ToDo 頁面
func RenderTodo(c *gin.Context) {
	c.HTML(http.StatusOK, "todo.html", nil)
}
