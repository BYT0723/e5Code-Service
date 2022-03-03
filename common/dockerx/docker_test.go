package dockerx

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/docker/docker/api/types"
	docker "github.com/docker/docker/client"
)

func TestFlag(t *testing.T) {
	cli, err := docker.NewClientWithOpts(docker.FromEnv)
	if err != nil {
		log.Fatal("Fail to new docker client, err: ", err.Error())
		return
	}

	// List Images
	images, _ := cli.ImageList(context.Background(), types.ImageListOptions{})
	fmt.Println("Images :")
	for i, v := range images {
		fmt.Printf("images-%d: %v\n", i, v)
	}

	// List Containers
	containers, _ := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})
	fmt.Println("Containers :")
	for i, v := range containers {
		fmt.Printf("container-%d: %v\n", i, v)
	}
}
