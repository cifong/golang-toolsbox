package router

import (
	"log"
)

func StartServer() {
	r := SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
