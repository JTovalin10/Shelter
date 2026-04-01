package models

type HostMetrics struct {
	// call Info() which returns *InfoStat
	Uptime uint64 `json:"uptime"`
	Procs uint64 `json:"procs`
	OS string `json:"os"`
}