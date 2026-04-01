package models

// Docs: https://pkg.go.dev/github.com/shirou/gopsutil/v4@v4.26.3

type Metrics struct {
	CPU CPUMetrics `json:"cpu"`
	Memory MemoryMetrics `json:"memory"`
	Disk DiskMetrics `json:"disk"`
	Network NetworkMetrics `json:"network"`
	Host HostMetrics `json:"host"`
}