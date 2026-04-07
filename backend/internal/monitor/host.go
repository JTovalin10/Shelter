package monitor

import (
	"shelter/backend/internal/models"

	"github.com/shirou/gopsutil/v4/host"
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

func CollectOS() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", err
	}
	return info.OS, nil
}

func CollectUptime() (uint64, error) {
	info, err := host.Info()
	if err != nil {
		return 0, err
	}
	return info.Uptime, nil
}

func CollectProcs() (uint64, error) {
	info, err := host.Info()
	if err != nil {
		return 0, err
	}
	return info.Procs, nil
}
