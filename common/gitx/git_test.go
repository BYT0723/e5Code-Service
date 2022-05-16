package gitx

import (
	"fmt"
	"log"
	"testing"

	git "github.com/go-git/go-git/v5"
)

func TestGit(t *testing.T) {
  cli := NewCliWithOpt("git.byt0723.xyz:22", "git", "wangtao")
  res, err := cli.ForkRegistry("BYT0723", "i3config", "git@gitee.com:wanngtao/configs.git")
  if err != nil {
    fmt.Println(res)
  }
}

func TestClone(t *testing.T) {
  if err := Clone("./tmp/git_registry/", "git@git.byt0723.xyz:test/test.git"); err != nil {
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
    }); err != nil {
    log.Fatal("Fail:", err.Error())
  }
  fmt.Println("Success")
}
