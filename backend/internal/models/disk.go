package models

type DiskMetrics struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
	Device      string  `json:"device"`
	ReadCount   uint64  `json:"readCount"`
	WriteCount  uint64  `json:"writeCount"`
}
