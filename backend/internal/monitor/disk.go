package monitor

import (
	"github.com/shirou/gopsutil/v4/disk"
	"shelter/backend/internal/models"
)

func CollectDisk() (models.DiskMetrics, err) {
	usage, err := disk.Usage("/") // gets *UsageStat, err from root

	if err != nil {
		return models.DiskMetrics{}, err
	}

	partitions, err := disk.Partitions(false) // returns physical devices only
	if err != nil {
		return models.DiskMetrics{}, err
	}

	counters, err := disk.IOCounters() // get all devices
	if err != nil {
		return models.DiskMetrics{}, err
	}

	var totalReads, totalWrites uint64
	for _, stats := range counters {
		totalReads += stats.ReadCount
		totalWrites += stats.WriteCount
	}

	results := {
		Total: usage.Total,
		Free: usage.Free,
		Used: usage.Used,
		UsedPercent: usage.UsedPercent,
		Device: partitions.Device,
		ReadCount: totalReads,
		WriteCount: WriteCount,
	}
	return results, nil
}