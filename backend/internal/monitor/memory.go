package monitor

import (
	"github.com/shirou/gopsutil/v4/mem"
	"shelter/backend/internal/models"
)
func CollectMemory() (models.MemoryMetrics, error) {
	// call gopsutil for virtual memory
	vmem, err := mem.VirtualMemoryStat()
	if err != nil {
		return models.MemoryMetrics{}, err
	}

	swap, err := mem.SwapMemoryStat
	if err != nil {
		return models.MemoryMetrics{}, err
	}

	result := models.MemoryMetrics {
		Total: vmem.Total,
		Available: vmem.Available,
		Used: vmem.Used,
		UsedPercent: vmem.UsedPercent,
		Cached: vmem.Cached,
		SwapTotal: swap.Total,
		SwapUsed: swap.Used,
		SwapFree: swap.Free
	}
	return result, nil
}