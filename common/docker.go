package common

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

const (
	Running = "running"
	Exit    = "exited"
)

type DockerClient struct {
	*client.Client
}

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

func (cli DockerClient) ListContainers(ctx context.Context, opt types.ContainerListOptions) (running []types.Container, exited []types.Container, err error) {
	cs, err := cli.ContainerList(ctx, opt)
	if err != nil {
		return nil, nil, err
	}
	for _, v := range cs {
		switch v.State {
		case Running:
			running = append(running, v)
		case Exit:
			exited = append(exited, v)
		}
	}
	return
}

func (cli DockerClient) ListStats(ctx context.Context, ids []string) (stats []*ContainerStats) {
	for i, v := range ids {
		rsp, err := cli.ContainerStats(ctx, v, false)
		if err != nil {
			log.Fatalf("Fail to get container stats(id: %s), err: %s", v, err.Error())
			continue
		}
		bs := make([]byte, 2048)
		len, _ := rsp.Body.Read(bs)
		stats = append(stats, &ContainerStats{})
		if err := json.Unmarshal(bs[:len], stats[i]); err != nil {
			log.Fatal("Fail to unmarshal stats, err: ", err.Error())
			continue
		}
	}
	return
}

func NewDockerClient(host, userID string) (*DockerClient, error) {
	dCli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithTLSClientConfig("./SSL/myDocker/ca.pem", "./SSL/myDocker/cert.pem", "./SSL/myDocker/key.pem"),
		client.WithHost(host),
	)
	if err != nil {
		return nil, err
	}
	return &DockerClient{dCli}, nil
}

func StoreActiveContainer(ctx context.Context, dockerHost string) error {
	// 新建docker client
	userid, err := GetUserID(ctx)
	if err != nil {
		return err
	}
	cli, err := NewDockerClient("https://42.192.5.238:2375", userid)
	if err != nil {
		return err
	}
	//获取当前docker daemon中所有的容器
	run, exit, err := cli.ListContainers(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return err
	}
	rIDs := []string{}
	eIDs := []string{}
	for _, v := range run {
		rIDs = append(rIDs, v.ID)
	}
	for _, v := range exit {
		eIDs = append(eIDs, v.ID)
	}

	stats := cli.ListStats(ctx, rIDs)

	config := InfluxConnConfig{
		Host: "http://frp.byt0723.xyz:8086",
		User: "root",
		Pass: "wangtao",
		DB:   "e5Code",
	}
	client, err := NewInfluxClient(config)
	if err != nil {
		return err
	}
	// Insert

	// tags := map[string]string{
	//     "io_read":         "read stream",
	//     "io_write":        "write stream",
	//     "cpu_total_usage": "%",
	//     "memory_usage":    "%",
	// }
	//
	for _, v := range stats {
		var (
			cpuUsage    float64 = 0
			memoryUsage float64 = 0
		)
		cpuDelta := float64(v.CPUStats.CPUUsage.TotalUsage) - float64(v.PrecpuStats.CPUUsage.TotalUsage)
		systemDelta := float64(v.CPUStats.SystemCpuUsage) - float64(v.PrecpuStats.SystemCpuUsage)
		if systemDelta > 0 {
			cpuUsage = cpuDelta / systemDelta * 100
		}
		if v.MemoryStats.Limit > 0 {
			memoryUsage = float64(v.MemoryStats.Usage) / float64(v.MemoryStats.Limit/100)
		}
		fields := map[string]interface{}{
			"id":              v.ID,
			"name":            v.Name[1:],
			"pid":             v.Pid.Current,
			"cpu_total_usage": cpuUsage,
			"cpu_nums":        v.CPUStats.OnlineCpus,
			"memory_usage":    memoryUsage,
			"block_in":        float64(v.BlkioStats.IOServiceBytesRecursive[0].Value) / 1024,
			"block_out":       float64(v.BlkioStats.IOServiceBytesRecursive[1].Value) / 1024,
			"net_in":          float64(v.Networks.Eth0.RXBytes) / 1024,
			"net_out":         float64(v.Networks.Eth0.TXBytes) / 1024,
		}
		if err := client.Insert("container_datas", nil, fields); err != nil {
			log.Fatalf("Fail to insert(name: %s), err: %s", v.Name, err.Error())
		}
	}
	return nil
}
