package monitor

import (
	"github.com/shirou/gopsutil/v4/cpu"
	"shelter/backend/internal/models"
)

func CollectCpu() (models.CPUMetrics, err) {
	overall, err := cpu.Percent(0, true) // returns per core
	
	if err != nil {
		return models.CPUMetrics{}, err
	}
	percore, err := cpu.Percent(0, false) // returns average
	if err != nil {
		return models.CPUMetrics{}, err
	}
	times, err := cpu.Times(false)
	if err != nil {
		return models.CPUMetrics{}, err
	}
	info, err := cpu.Info()
	if err != nil {
		return models.CPUMetrics{}, err
	}
	result := models.CPUMetrics{
		OverallPercent: overall[0],
		PerCorePercent: perCore,
		User:           times[0].User,
		System:         times[0].System,
		Idle:           times[0].Idle,
		Iowait:         times[0].Iowait,
		Steal:          times[0].Steal,
		ModelName:      info[0].ModelName,
		Cores:          info[0].Cores,
		Mhz:            info[0].Mhz,
		CacheSize:      info[0].CacheSize,
		Timestamp:      time.Now().UnixMilli(),
  	}
	return result
}