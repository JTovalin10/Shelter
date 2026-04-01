package monitor

import (
	"github.com/shirou/gopsutil/v4/net"
	"shelter/backend/internal/models"
)

func CollectNetwork() (models.NetworkMetrics, err) {
	IO, err := net.IOCounters(false) // returns sum of all information
	if err != nil {
		return models.NetworkMetrics, err
	}

	addr, err := net.Addr
	if err != nil {
		return models.NetworkMetrics{}, err
	}

	result := {
		Name: IO.Name,
		BytesSent: IO.BytesSent,
		BytesRecv: IO.BytesRecv,
		PacketsSent: IO.BytesSent,
		PacketsRecv: IO.BytesRecv,
		IP: addr.IP,
		Port: addr.Port
	}

	return result, nil
}