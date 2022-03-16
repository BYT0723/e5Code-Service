package dockerx

import "time"

type ContainerStats struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Read        time.Time   `json:"read"`
	Pid         PID         `json:"pid"`
	BlkioStats  BlkioStats  `json:"blkio_stats"`
	CPUStats    CPUStats    `json:"cpu_stats"`
	PrecpuStats CPUStats    `json:"precpu_stats"`
	MemoryStats MemoryStats `json:"memory_stats"`
	Networks    Networks    `json:"networks"`
}

type PID struct {
	Current int64 `json:"current"`
	Limit   int64 `json:"limit"`
}

type BlkioStats struct {
	IOServiceBytesRecursive []IOServiceBytesRecursive `json:"io_service_bytes_recursive"`
}

type IOServiceBytesRecursive struct {
	Major int64  `json:"major"`
	Minor int64  `json:"minor"`
	Op    string `json:"op"`
	Value int64  `json:"value"`
}

type CPUStats struct {
	SystemCpuUsage int64    `json:"system_cpu_usage"`
	OnlineCpus     int64    `json:"online_cpus"`
	CPUUsage       CPUUsage `json:"cpu_usage"`
}

type CPUUsage struct {
	TotalUsage        int64 `json:"total_usage"`
	UsageInKernelmode int64 `json:"usage_in_kernelmode"`
	UsageInUsermode   int64 `json:"usage_in_usermode"`
}

type MemoryStats struct {
	Usage int64 `json:"usage"`
	Limit int64 `json:"limit"`
}

type Networks struct {
	Eth0 Eth0 `json:"eth0"`
}

type Eth0 struct {
	RXBytes   int64 `json:"rx_bytes"`
	RXPackets int64 `json:"rx_packets"`
	RXErrors  int64 `json:"rx_errors"`
	RXDropped int64 `json:"rx_dropped"`
	TXBytes   int64 `json:"tx_bytes"`
	TXPackets int64 `json:"tx_packets"`
	TXErrors  int64 `json:"tx_errors"`
	TXDropped int64 `json:"tx_dropped"`
}
