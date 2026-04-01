package handlers

// Docs: https://pkg.go.dev/github.com/gorilla/websocket
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

}

// collects all metrics into one SystemMetriics struct
func collectAll() models.SystemMetrics {

}