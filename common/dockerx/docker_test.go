package dockerx

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/zeromicro/go-zero/core/logx"
)

func TestFlag(t *testing.T) {
	ctx := context.Background()
	cli, err := NewDockerClient()
	if err != nil {
		log.Fatal("Fail to new docker client, err: ", err.Error())
		return
	}
	stream, err := cli.BuildImage(ctx, "/home/tao/e5Code-Service.tar", types.ImageBuildOptions{
		Tags:       []string{"test"},
		Dockerfile: "./service/user/api/Dockerfile",
	})
	if err != nil {
		log.Fatal("Fail to Build Image:", err.Error())
		return
	}
	go func(stream *bufio.Reader) {
		logFile, _ := os.OpenFile("./test.log", os.O_RDWR|os.O_CREATE, 0666)
		stream.WriteTo(logFile)
	}(stream)
	for {
		data, _, err := stream.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Printf("data: %v\n", string(data))
	}
}

func TestPull(t *testing.T) {
	cli, err := NewDockerClient()
	if err != nil {
		log.Fatal("Fail to new docker client, err: ", err.Error())
		return
	}
	reader, err := cli.ImagePull(context.Background(), "alpine", types.ImagePullOptions{})
	if err != nil {
		logx.Error(err)
		return
	}
	stream := bufio.NewReader(reader)
	for {
		line, _, err := stream.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Printf("line: %v\n", string(line))
	}
}

func TestRun(t *testing.T) {
	cli, err := NewDockerClient()
	if err != nil {
		log.Fatal("Fail to new docker client, err: ", err.Error())
		return
	}

	cli.ContainerCreate(context.Background(), &container.Config{
		Image: "redis:alpine",
		Tty:   false,
	}, nil, nil, nil, "redis-test")
}

func TestScan(t *testing.T) {
	// imageID := ""
	// str := `{"stream":"Successfully built ad1e87655d3a\n"}`
	str := ""
	reg1 := regexp.MustCompile(`built ([a-z0-9]+)`)
	res := strings.Split(reg1.FindString(str), " ")[1]
	fmt.Printf("res: %v\n", res)
}
