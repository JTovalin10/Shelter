package models

type MemoryMetrics struct {
	// virtual memory
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
	Cached      uint64  `json:"cached"`

	// swap
	SwapTotal uint64 `json:"swapTotal"`
	SwapUsed  uint64 `json:"swapUsed"`
	SwapFree  uint64 `json:"swapFree"`
}
