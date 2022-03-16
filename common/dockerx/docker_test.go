package dockerx

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/docker/docker/api/types"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
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

func TestClone(t *testing.T) {
	auth, err := ssh.NewPublicKeysFromFile("git", "/home/tao/.ssh/id_rsa", "")
	if err != nil {
		log.Fatal("Fail to get ssh public key: ", err.Error())
		return
	}

	if err := Clone(GitCloneOpt{
		Local: "./tmp/git_registry/",
		CloneOptions: &git.CloneOptions{
			URL:  "git@github.com:BYT0723/e5Code.git",
			Auth: auth,
		},
	}); err != nil {
		log.Fatal("Fail to Clone registry: ", err.Error())
		return
	}
}

func TestTar(t *testing.T) {
	if err := TarProject("/home/tao/e5Code-Service.tar", "/home/tao/Documents/Github/e5Code-Service"); err != nil {
		log.Fatal(err)
		return
	}
}
