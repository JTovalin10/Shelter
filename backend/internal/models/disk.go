package models

type DiskMetrics struct {
	// From UsageStat struct
	// call disk.Usage("/") returns *UsageStat
	Free uint64 `json:"Free"`
	Used uint64 `json:"Used"`
	UsedPercent float64 `json:"UsedPercent"`

	// from PartitionState struct
	// call disk.Partitions(false) returns []PartitionStat
	Device string `json:"Device"`

	// From IOCounterState struct
	// call disk.IOCounterts returns map[string]IOCounterStat
	ReadCount uint64 `json:"readCount"`
	WriteCount uint64 `json:"writeCount"`

}