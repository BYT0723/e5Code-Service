package gitx

import (
	git "github.com/go-git/go-git/v5"
)

type GitCloneOpt struct {
	Local  string
	IsBare bool
	*git.CloneOptions
}

// Clone远程仓库到本地，返回Clone到的本地URL
func Clone(opt GitCloneOpt) error {
	if _, err := git.PlainClone(opt.Local, opt.IsBare, opt.CloneOptions); err != nil {
		return err
	}
	return nil
}
