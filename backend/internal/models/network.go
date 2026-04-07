package models

type NetworkMetrics struct {
	BytesSent uint64 `json:"bytesSent"` // bytes per second
	BytesRecv uint64 `json:"bytesRecv"` // bytes per second
	IP        string `json:"ip"`
	MAC       string `json:"mac"`
}
