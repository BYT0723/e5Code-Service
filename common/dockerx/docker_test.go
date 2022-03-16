package dockerx

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/docker/docker/api/types"
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
	for true {
		line, err := stream.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal("Fail to Read Result: ", err.Error())
			return
		}
		fmt.Println(line)
	}
}

func TestTar(t *testing.T) {
	if err := TarProject("/home/tao/e5Code-Service.tar", "/home/tao/Documents/Github/e5Code-Service"); err != nil {
		log.Fatal(err)
		return
	}
}
