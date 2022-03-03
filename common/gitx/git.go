package gitx

import (
	"fmt"
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/zeromicro/go-zero/core/logx"
)

func Clone(url string) {
	// res, err := git.Clone(url)
	// if err != nil {
	//     logx.Error("Fail to clone : ", err.Error())
	//     return
	// }
	// fmt.Printf("res: %v\n", res)

	rsp, err := git.PlainClone("/tmp/gitCloneProjectName", false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		logx.Error("Fail to PlainClone Registry: ", err.Error())
		return
	}
	fmt.Printf("rsp: %v\n", rsp)
}
