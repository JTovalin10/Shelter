package handlers

Docs: https://pkg.go.dev/github.com/gorilla/websocket
import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/websocket"
	"shelter/backend/internal/monitor"
)

// the upgrader config (package-level var, not a function)
var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024
}

// HTTP handler that upgrades to WebSocket and starts streaming
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	con, err := upgrader.upgrader(w, r, nil)
	if err != nil {
		return
	}
	// closes when function returns
	defer con.close()

	ticker := time.newTicker(2 * time.Second)

	for range ticker.C {
		metrics := collectAll()
		data, err := json.Marshal(metrics)
		if err != nil {
			return // issue with json
		}
		if err := con.WriteMessage(websocket.TextMessage, data); err != nil {
			return // client disconnected
		}
	}
}

// collects all metrics into one SystemMetriics struct
func collectAll() models.SystemMetrics {
	cpu, _ = CollectCpu()
	disk, _ = CollectDisk()
	host, _ = CollectHost()
	mem, _= CollectMemory()
	net, _ = CollectNetwork()
	return models.SystemMetrics{CPU: cpu, Memory: mem, Disk: disk, Network: net, Host: host}
}