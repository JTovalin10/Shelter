package handlers

// Docs: https://pkg.go.dev/github.com/gorilla/websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"shelter/backend/internal/models"
	"shelter/backend/internal/monitor"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // restrict in production
	},
}

func RegisterRoutes(r chi.Router) {
	r.Get("/ws", handleWebSocket)
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}
	defer conn.Close()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		metrics := collectAll()
		data, err := json.Marshal(metrics)
		if err != nil {
			return
		}
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			return // client disconnected
		}
	}
}

func collectAll() models.Metrics {
	cpu, _ := monitor.CollectCpu()
	disk, _ := monitor.CollectDisk()
	host, _ := monitor.CollectHost()
	mem, _ := monitor.CollectMemory()
	net, _ := monitor.CollectNetwork()
	return models.Metrics{CPU: cpu, Memory: mem, Disk: disk, Network: net, Host: host}
}
