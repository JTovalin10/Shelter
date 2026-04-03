package monitor

import (
	"github.com/shirou/gopsutil/v4/mem"
	"shelter/backend/internal/models"
)

func CollectMemory() (models.MemoryMetrics, error) {
	vmem, err := mem.VirtualMemory()
	if err != nil {
		return models.MemoryMetrics{}, err
	}

	swap, err := mem.SwapMemory()
	if err != nil {
		return models.MemoryMetrics{}, err
	}

	result := models.MemoryMetrics{
		Total:       vmem.Total,
		Available:   vmem.Available,
		Used:        vmem.Used,
		UsedPercent: vmem.UsedPercent,
		Cached:      vmem.Cached,
		SwapTotal:   swap.Total,
		SwapUsed:    swap.Used,
		SwapFree:    swap.Free,
	}
	return result, nil
}
