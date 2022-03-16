package dockerx

import (
	"bufio"
	"context"
	"encoding/json"
	"os"
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

func NewDockerClient() (*DockerClient, error) {
	dCli, err := client.NewClientWithOpts(
		client.FromEnv,
	)
	if err != nil {
		return nil, err
	}
	return &DockerClient{dCli}, nil
}

func (cli *DockerClient) ListContainers(ctx context.Context, opt types.ContainerListOptions) (running []types.Container, exited []types.Container, err error) {
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

func (cli *DockerClient) ListStats(ctx context.Context, ids []string) (stats []*ContainerStats) {
	for i, v := range ids {
		rsp, err := cli.ContainerStats(ctx, v, false)
		if err != nil {
			continue
		}
		bs := make([]byte, 2048)
		len, _ := rsp.Body.Read(bs)
		stats = append(stats, &ContainerStats{})
		if err := json.Unmarshal(bs[:len], stats[i]); err != nil {
			continue
		}
	}
	return
}

// 构建镜像 传入项目的tar压缩文件，以及opt中的Tag，Dockerfile等参数
func (cli *DockerClient) BuildImage(ctx context.Context, tarFile string, opt types.ImageBuildOptions) (*bufio.Reader, error) {
	dir, err := os.Open(tarFile)
	if err != nil {
		return nil, err
	}
	rsp, err := cli.ImageBuild(ctx, dir, opt)
	if err != nil {
		return nil, err
	}
	return bufio.NewReader(rsp.Body), nil
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
