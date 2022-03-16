package gitx

import (
	"fmt"
	"testing"
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
