package models

type HostMetrics struct {
	Uptime uint64 `json:"uptime"`
	Procs  uint64 `json:"procs"`
	OS     string `json:"os"`
}
