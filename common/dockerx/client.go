package dockerx

import (
	"bufio"
	"context"
	"e5Code-Service/service/cd/model"
	"encoding/base64"
	"encoding/json"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
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
	defer os.Remove(tarFile)

	rsp, err := cli.ImageBuild(ctx, dir, opt)
	if err != nil {
		return nil, err
	}
	return bufio.NewReader(rsp.Body), nil
}

// 上传镜像
func (cli *DockerClient) UploadImage(ctx context.Context, imageName string, auth types.AuthConfig) (*bufio.Reader, error) {
	// // 给镜像打Tag
	// if err := cli.ImageTag(ctx, imageName, fmt.Sprintf("%s/%s/%s", host, username, imageName)); err != nil {
	//   return nil, err
	// }

	// Auth鉴权
	encodeJson, _ := json.Marshal(auth)
	authStr := base64.URLEncoding.EncodeToString(encodeJson)

	// 镜像推送
	reader, err := cli.ImagePush(ctx, imageName, types.ImagePushOptions{
		RegistryAuth: authStr,
	})
	if err != nil {
		return nil, err
	}
	res := bufio.NewReader(reader)
	return res, nil
}

type RunOption struct {
	ImageName     string
	ContainerName string
	Hostname      string
	Envs          model.Envs
	Ports         model.PortList
	RestartPolicy *model.RestartPolicy
	Architecture  string
	OS            string
}

func (cli *DockerClient) Run(ctx context.Context, opt RunOption) (id string, runLog *bufio.Reader, err error) {
	// 拉取镜像
	if _, err = cli.ImagePull(ctx, opt.ImageName, types.ImagePullOptions{}); err != nil {
		return
	}

	// 端口映射处理
	portMap := make(nat.PortMap)
	for _, pb := range opt.Ports {
		port, _ := nat.NewPort(pb.Proto, pb.Port)
		portMap[port] = []nat.PortBinding{
			{
				HostIP:   pb.HostIP,
				HostPort: pb.HostPort,
			},
		}
	}

	// 容器配置
	containerConfig := &container.Config{
		Image:    opt.ImageName,
		Hostname: opt.Hostname,
		Env:      opt.Envs,
	}

	// 容器host配置
	hostConfig := &container.HostConfig{
		// 端口绑定
		PortBindings: portMap,
		// 重启策略
		RestartPolicy: container.RestartPolicy{
			Name:              opt.RestartPolicy.Name,
			MaximumRetryCount: opt.RestartPolicy.MaximumRetryCount,
		},
		// 自动移除同名容器
		AutoRemove: true,
	}
	// 容器网络配置
	networkConfig := &network.NetworkingConfig{}

	// 容器平台配置
	platformConfig := &v1.Platform{
		// cpu架构 amd64/arm等
		Architecture: opt.Architecture,
		// 系统类型 linux/windows等
		OS: opt.OS,
	}

	// 创建容器
	createRsp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, networkConfig, platformConfig, opt.ContainerName)
	if err != nil {
		return
	}

	// 启动容器
	if err = cli.ContainerStart(ctx, createRsp.ID, types.ContainerStartOptions{}); err != nil {
		return
	}

	// 获取容器日志
	logsReader, err := cli.ContainerLogs(ctx, opt.ContainerName, types.ContainerLogsOptions{})
	if err != nil {
		return
	}

	return createRsp.ID, bufio.NewReader(logsReader), nil
}
