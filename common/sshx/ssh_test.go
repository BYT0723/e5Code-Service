package sshx

import (
	"fmt"
	"testing"
)

func TestSSH(t *testing.T) {
	cli := NewCli("git.byt0723.xyz:22", "git", "wangtao")
	cli.Run("mkdir test")
	fmt.Printf("cli.LastResult: %v\n", cli.LastResult)
}
