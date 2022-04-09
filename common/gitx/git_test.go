package gitx

import (
	"fmt"
	"log"
	"testing"

	"github.com/go-git/go-git/plumbing/transport/ssh"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
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

func TestList(t *testing.T) {
	rep, _ := git.PlainOpen("/home/tao/Documents/Github/e5code")
	result, _ := ListFile(rep, "", true, false)
	for _, v := range result {
		fmt.Printf("v: %v\n", v.CommitInfo)
		if v.Children != nil {
			for _, v2 := range v.Children {
				fmt.Printf("v2: %v\n", v2.CommitInfo)
			}
		}
	}
}

func TestUpdateFile(t *testing.T) {
	rep, _ := git.PlainOpen("/home/tao/Documents/Gitee/configs")
	body, _ := ReadFile(rep, "README", true)
	fmt.Printf("body: %v\n", body)

	if err := UpdateFile(rep, "README", []byte("Test")); err != nil {
		log.Fatal("Fail to updateFile:", err.Error())
		return
	}
	body, _ = ReadFile(rep, "README", false)
	fmt.Printf("body: %v\n", body)
}

func TestGitStauts(t *testing.T) {
	// rep, _ := git.PlainOpen("/home/tao/Documents/Gitee/configs")
}

func TestPush(t *testing.T) {
	rep, _ := git.PlainOpen("/home/tao/Documents/Gitee/configs")
	if err := Commit(rep, &CommitOption{
		Msg:    "go-git测试",
		Author: "wangtao",
		Email:  "1151713064@qq.com",
		Remote: "origin",
		BasicAuth: &http.BasicAuth{
			Username: "13164884812",
			Password: "WTlove0910..",
		},
	}); err != nil {
		log.Fatal("Fail:", err.Error())
	}
	fmt.Println("Success")
}
