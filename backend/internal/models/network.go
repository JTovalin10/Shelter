package models

type NetworkMetrics struct {
	// IOCountersStat
	// call n.IOCounterStat which returns a string
	Name        string `json:"name"`        // interface name
	BytesSent   uint64 `json:"bytesSent"`   // number of bytes sent
	BytesRecv   uint64 `json:"bytesRecv"`   // number of bytes received
	PacketsSent uint64 `json:"packetsSent"` // number of packets sent
	PacketsRecv uint64 `json:"packetsRecv"` // number of packets received

	// Addr
	// call a.Addr which returns a string
	IP string `json:"ip"`
	Port uint32 `json:"port"`
}