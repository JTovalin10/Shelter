package monitor

import (
	"github.com/shirou/gopsutil/v4/disk"
	"shelter/backend/internal/models"
)

func CollectDisk() (models.DiskMetrics, error) {
	usage, err := disk.Usage("/")
	if err != nil {
		return models.DiskMetrics{}, err
	}

	partitions, err := disk.Partitions(false)
	if err != nil {
		return models.DiskMetrics{}, err
	}

	counters, err := disk.IOCounters()
	if err != nil {
		return models.DiskMetrics{}, err
	}

	var totalReads, totalWrites uint64
	for _, stats := range counters {
		totalReads += stats.ReadCount
		totalWrites += stats.WriteCount
	}

	result := models.DiskMetrics{
		Total:       usage.Total,
		Free:        usage.Free,
		Used:        usage.Used,
		UsedPercent: usage.UsedPercent,
		Device:      partitions[0].Device,
		ReadCount:   totalReads,
		WriteCount:  totalWrites,
	}
	return result, nil
}
