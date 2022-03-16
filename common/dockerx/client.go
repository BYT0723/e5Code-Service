package dockerx

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

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

// 上传镜像
func (cli *DockerClient) UploadImage(ctx context.Context, host, username, imageName string) (*bufio.Reader, error) {
	// 给镜像打Tag
	if err := cli.ImageTag(ctx, imageName, fmt.Sprintf("%s/%s/%s", host, username, imageName)); err != nil {
		return nil, err
	}

	// Auth鉴权
	encodeJson, _ := json.Marshal(types.AuthConfig{})
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
