package dockerx

import (
	"archive/tar"
	"context"
	"e5Code-Service/common/influxx"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/docker/api/types"
)

// 存储docker中所有容器的数据
func StoreActiveContainer(ctx context.Context, dockerHost string) error {
	// 新建docker client
	cli, err := NewDockerClient()
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

	config := influxx.InfluxConnConfig{
		Host: "http://frp.byt0723.xyz:8086",
		User: "root",
		Pass: "wangtao",
		DB:   "e5Code",
	}
	client, err := influxx.NewInfluxClient(config)
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

// 打包指定的文件路径
// target : tar包路径
// source : 源文件夹路径
func TarProject(target, source string) error {
	tf, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer tf.Close()

	w := tar.NewWriter(tf)
	defer w.Close()

	if err := filepath.Walk(source, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}

		hr := &tar.Header{
			Name: strings.Replace(path, source, ".", -1),
			Size: info.Size(),
			Mode: 0666,
		}

		w.WriteHeader(hr)
		var buf [1024]byte
		for {
			n, err := f.Read(buf[:])
			w.Write(buf[:n])
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
