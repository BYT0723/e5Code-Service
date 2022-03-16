package gitx

import (
	"fmt"
	"log"
	"testing"

	"github.com/go-git/go-git/plumbing/transport/ssh"
	git "github.com/go-git/go-git/v5"
)

func TestGit(t *testing.T) {
	cli := NewCliWithOpt("git.byt0723.xyz:22", "git", "wangtao")
	res, err := cli.CreateUser("test")
	fmt.Printf("res: %v\n", res)
	fmt.Printf("err: %v\n", err)
	if err != nil {
		return
	}

	res, err = cli.CreateRegistry("test", "testFirst")
	fmt.Printf("res: %v\n", res)
	fmt.Printf("err: %v\n", err)
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
