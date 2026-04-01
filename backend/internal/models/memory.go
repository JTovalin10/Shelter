package models

type MemoryMetrics struct {
	// virtual memory
	// call m.VirtualMemoryStat which returns a string
	Total uint64 `json:"total"`
	Available uint64 `json:"available"`
	Used uint64 `json:"used"`
	UsedPercent float64 `json:"usedPercent"`\
	// Linux specific numbers
	Cached uint64 `json:"cached"`

	// swap
	// call m.SwapMemoryState() which returns a string
	SwapTotal uint64 `json:"SwapTotal`
	SwapUsed uint64 `json:"SwapUsed"`
	SwapFree uint64 `json:"SwapFree"`
}