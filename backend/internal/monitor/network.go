package monitor

import (
	gopsnet "github.com/shirou/gopsutil/v4/net"
	"shelter/backend/internal/models"
)

func CollectNetwork() (models.NetworkMetrics, error) {
	io, err := gopsnet.IOCounters(false) // false = aggregate all interfaces
	if err != nil || len(io) == 0 {
		return models.NetworkMetrics{}, err
	}

	interfaces, err := gopsnet.Interfaces()
	if err != nil {
		return models.NetworkMetrics{}, err
	}

	var ip string
	for _, iface := range interfaces {
		if iface.Name != "lo" && len(iface.Addrs) > 0 {
			ip = iface.Addrs[0].Addr
			break
		}
	}

	result := models.NetworkMetrics{
		Name:        io[0].Name,
		BytesSent:   io[0].BytesSent,
		BytesRecv:   io[0].BytesRecv,
		PacketsSent: io[0].PacketsSent,
		PacketsRecv: io[0].PacketsRecv,
		IP:          ip,
	}
	return result, nil
}
