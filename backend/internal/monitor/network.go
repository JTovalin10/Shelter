package monitor

import (
	"strings"
	"time"

	gopsnet "github.com/shirou/gopsutil/v4/net"
	"shelter/backend/internal/models"
)

type Bytes struct {
	BytesSent uint64
	BytesRecv uint64
}

func CollectNetwork() (models.NetworkMetrics, error) {
	bytes, err := getBandwidth()
	if err != nil {
		return models.NetworkMetrics{}, err
	}

	interfaces, err := gopsnet.Interfaces()
	if err != nil {
		return models.NetworkMetrics{}, err
	}

	var ip string
	for _, iface := range interfaces {
		if strings.HasPrefix(iface.Name, "lo") {
			continue
		}
		for _, addr := range iface.Addrs {
			raw := strings.Split(addr.Addr, "/")[0]
			if !strings.Contains(raw, ":") { // skip IPv6
				ip = raw
				break
			}
		}
		if ip != "" {
			break
		}
	}

	result := models.NetworkMetrics{
		BytesSent: bytes.BytesSent,
		BytesRecv: bytes.BytesRecv,
		IP:        ip,
	}
	return result, nil
}

func CollectMAC() (string, error) {
	interfaces, err := gopsnet.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range interfaces {
		if strings.HasPrefix(iface.Name, "lo") {
			continue
		}
		if iface.HardwareAddr != "" {
			return iface.HardwareAddr, nil
		}
	}
	return "", nil
}

func getBandwidth() (Bytes, error) {
	counter1, err := gopsnet.IOCounters(false)
	if err != nil {
		return Bytes{}, err
	}

	t1 := time.Now()
	time.Sleep(1 * time.Second)

	counter2, err := gopsnet.IOCounters(false)
	if err != nil {
		return Bytes{}, err
	}

	elapsed := time.Since(t1).Seconds()

	var sent, recv uint64
	length := len(counter2)
	for i, c2 := range counter2 {
		c1 := counter1[i]
		sent += uint64(float64(c2.BytesSent-c1.BytesSent) / elapsed)
		recv += uint64(float64(c2.BytesRecv-c1.BytesRecv) / elapsed)
	}

	if length == 0 {
		return Bytes{}, nil
	}

	return Bytes{
		BytesSent: sent / uint64(length),
		BytesRecv: recv / uint64(length),
	}, nil
}
