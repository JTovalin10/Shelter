package handlers

// Docs: https://pkg.go.dev/github.com/gorilla/websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"shelter/backend/internal/models"
	"shelter/backend/internal/monitor"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"

	"github.com/sourcegraph/conc"
)

const CACHE_TTL time.Duration = 5 * time.Minute

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
	// set up
	defer conn.Close()
	start := time.Now()
	os, _ := monitor.CollectOS()
	mac, _ := monitor.CollectMAC()
	uptime, _ := monitor.CollectUptime()
	ticker := time.NewTicker(2 * time.Second)
	uptimeTicker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	defer uptimeTicker.Stop()

	// send initial data immediately
	metrics := collectAll(os, mac, uptime)
	sendData(metrics, conn)
	// loop every 2 seconds and send data
	for {
		select {
		case <-uptimeTicker.C:
			uptime, _ = monitor.CollectUptime()

		case <-ticker.C:
			metrics := models.Metrics{}
			if time.Since(start) >= CACHE_TTL {
				start = time.Now()
				// refresh cache
				metrics = collectAll(os, mac, uptime)
			} else {
				metrics = collect_non_static(os, mac, uptime)
			}
			var failed bool = sendData(metrics, conn)
			if failed {
				return
			}
		}
	}
}

func sendData(metrics models.Metrics, conn *websocket.Conn) bool {
	data, err := json.Marshal(metrics)
	if err != nil {
		return true
	}
	if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
		return true
	}
	return false
}

func collectAll(os string, mac string, uptime uint64) models.Metrics {
	var cpu models.CPUMetrics
	var mem models.MemoryMetrics
	var net models.NetworkMetrics
	var disk models.DiskMetrics
	var procs uint64

	wg := conc.NewWaitGroup()
	wg.Go(func() { cpu, _ = monitor.CollectCpu() })
	wg.Go(func() { mem, _ = monitor.CollectMemory() })
	wg.Go(func() { net, _ = monitor.CollectNetwork() })
	wg.Go(func() { disk, _ = monitor.CollectDisk() })
	wg.Go(func() { procs, _ = monitor.CollectProcs() })
	wg.Wait()

	net.MAC = mac
	host := models.HostMetrics{OS: os, Uptime: uptime, Procs: procs}
	return models.Metrics{CPU: cpu, Memory: mem, Disk: disk, Network: net, Host: host}
}

func collect_non_static(os string, mac string, uptime uint64) models.Metrics {
	var cpu models.CPUMetrics
	var mem models.MemoryMetrics
	var net models.NetworkMetrics
	var procs uint64

	wg := conc.NewWaitGroup()
	wg.Go(func() { cpu, _ = monitor.CollectCpu() })
	wg.Go(func() { mem, _ = monitor.CollectMemory() })
	wg.Go(func() { net, _ = monitor.CollectNetwork() })
	wg.Go(func() { procs, _ = monitor.CollectProcs() })
	wg.Wait()

	net.MAC = mac
	host := models.HostMetrics{OS: os, Uptime: uptime, Procs: procs}
	return models.Metrics{CPU: cpu, Memory: mem, Network: net, Host: host}
}
