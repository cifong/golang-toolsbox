package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cifong/golang-toolsbox/internal/system"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		allowedOrigins := map[string]bool{
			"http://localhost:8080": true,
			"http://127.0.0.1:8080": true,
		}
		fmt.Println("🔍 Origin:", r.Header.Get("Origin"))
		return allowedOrigins[origin]
	},
}

func GetSystemInfoWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket 升級失敗:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "WebSocket upgrade failed"})
		return
	}
	defer conn.Close()

	for {
		info, err := system.GetSystemInfo()
		if err != nil {
			errMsg := gin.H{"error": err.Error()}
			msg, _ := json.Marshal(errMsg)
			conn.WriteMessage(websocket.TextMessage, msg)
			break
		}

		msg, err := json.Marshal(info)
		if err != nil {
			log.Println("JSON 編碼失敗:", err)
			break
		}

		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("傳送資料失敗:", err)
			break
		}

		time.Sleep(2 * time.Second)
	}
}
