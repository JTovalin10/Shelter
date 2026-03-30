# Backend System Metrics Design

## Overview

Go backend that collects system metrics (CPU, memory, disk, network, host) using gopsutil v4 and streams them to the frontend over WebSocket. Each metric has its own collector file and model file.

## Dependencies

- `github.com/go-chi/chi/v5` — router
- `github.com/go-chi/cors` — CORS middleware
- `github.com/gorilla/websocket` — WebSocket
- `github.com/shirou/gopsutil/v4` — system metrics

## File Structure

```
backend/
├── cmd/server/main.go          — HTTP server, chi router, CORS, WebSocket endpoint
├── internal/
│   ├── handlers/
│   │   └── ws.go               — WebSocket upgrade handler, read/write loop
│   ├── models/
│   │   ├── cpu.go              — CPUMetrics struct
│   │   ├── memory.go           — MemoryMetrics struct
│   │   ├── disk.go             — DiskMetrics struct
│   │   ├── network.go          — NetworkMetrics struct
│   │   ├── host.go             — HostMetrics struct
│   │   └── metrics.go          — SystemMetrics aggregate struct
│   └── monitor/
│       ├── cpu.go              — CollectCPU()
│       ├── memory.go           — CollectMemory()
│       ├── disk.go             — CollectDisk()
│       ├── network.go          — CollectNetwork()
│       └── host.go             — CollectHost()
```

## Models

### CPUMetrics

```go
type CPUCore struct {
    Core    int     `json:"core"`
    Usage   float64 `json:"usage"`
}

type CPUMetrics struct {
    Cores   []CPUCore `json:"cores"`
    Overall float64   `json:"overall"`
}
```

### MemoryMetrics

```go
type MemoryMetrics struct {
    Total     uint64  `json:"total"`
    Used      uint64  `json:"used"`
    Available uint64  `json:"available"`
    UsedPct   float64 `json:"used_pct"`
    SwapTotal uint64  `json:"swap_total"`
    SwapUsed  uint64  `json:"swap_used"`
    SwapPct   float64 `json:"swap_pct"`
}
```

### DiskMetrics

```go
type DiskPartition struct {
    Device     string  `json:"device"`
    Mountpoint string  `json:"mountpoint"`
    Fstype     string  `json:"fstype"`
    Total      uint64  `json:"total"`
    Used       uint64  `json:"used"`
    Free       uint64  `json:"free"`
    UsedPct    float64 `json:"used_pct"`
}

type DiskIO struct {
    Device     string `json:"device"`
    ReadBytes  uint64 `json:"read_bytes"`
    WriteBytes uint64 `json:"write_bytes"`
}

type DiskMetrics struct {
    Partitions []DiskPartition `json:"partitions"`
    IO         []DiskIO        `json:"io"`
}
```

### NetworkMetrics

```go
type NetworkInterface struct {
    Name      string `json:"name"`
    BytesSent uint64 `json:"bytes_sent"`
    BytesRecv uint64 `json:"bytes_recv"`
    PktsSent  uint64 `json:"pkts_sent"`
    PktsRecv  uint64 `json:"pkts_recv"`
}

type NetworkMetrics struct {
    Interfaces []NetworkInterface `json:"interfaces"`
}
```

### HostMetrics

```go
type HostMetrics struct {
    Hostname string `json:"hostname"`
    Uptime   uint64 `json:"uptime"`
    OS       string `json:"os"`
    Platform string `json:"platform"`
    Version  string `json:"version"`
    Arch     string `json:"arch"`
}
```

### SystemMetrics (aggregate)

```go
type SystemMetrics struct {
    CPU     CPUMetrics     `json:"cpu"`
    Memory  MemoryMetrics  `json:"memory"`
    Disk    DiskMetrics    `json:"disk"`
    Network NetworkMetrics `json:"network"`
    Host    HostMetrics    `json:"host"`
}
```

## Collectors

Each collector is a standalone function in `monitor/`:

- `CollectCPU() (models.CPUMetrics, error)` — calls `cpu.Percent()` per-core and overall
- `CollectMemory() (models.MemoryMetrics, error)` — calls `mem.VirtualMemory()` and `mem.SwapMemory()`
- `CollectDisk() (models.DiskMetrics, error)` — calls `disk.Partitions()`, `disk.Usage()`, `disk.IOCounters()`
- `CollectNetwork() (models.NetworkMetrics, error)` — calls `net.IOCounters()`
- `CollectHost() (models.HostMetrics, error)` — calls `host.Info()`

## WebSocket Handler

- Endpoint: `GET /ws`
- On connect: upgrade to WebSocket via gorilla/websocket
- Server sends a JSON `SystemMetrics` payload every 2 seconds
- On client disconnect: clean up goroutine
- Upgrader allows all origins in dev (CORS handled at HTTP level)

## Server (main.go)

- Listen on `:8080`
- Chi router with CORS middleware (allow `localhost:5173` for Vite dev server)
- Single route: `GET /ws` -> WebSocket handler
- Graceful shutdown on SIGINT/SIGTERM

## Error Handling

- Individual collector errors are logged but don't crash the server
- If a collector fails, its field is omitted or zero-valued in the payload
- WebSocket write errors close the connection and exit the goroutine
