package monitor

import (
	"github.com/shirou/gopsutil/v4/host"
	"shelter/backend/internal/models"
)

func CollectHost() (models.HostMetrics, error) {
	info, err := host.Info()
	if err != nil {
		return models.HostMetrics{}, err
	}

	result := models.HostMetrics{
		Uptime: info.Uptime,
		Procs:  info.Procs,
		OS:     info.OS,
	}
	return result, nil
}
