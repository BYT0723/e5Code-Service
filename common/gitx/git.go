package gitx

import (
	git "github.com/go-git/go-git/v5"
)

type File struct {
	Name   string
	IsFile bool
	Size   int64 //unit:Byte
	Path   string
}

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

// 拉取远程仓库最新提交
func Pull(rep *git.Repository, remoteName string) error {
	w, err := rep.Worktree()
	if err != nil {
		return err
	}
	return w.Pull(&git.PullOptions{RemoteName: remoteName})
}

// 获取指定Path下Files
func ListFile(rep *git.Repository, path string) (files []*File, err error) {
	// 获取Head指针
	ref, err := rep.Head()
	if err != nil {
		return
	}

	// 获取Head的Hash的提交
	commit, err := rep.CommitObject(ref.Hash())
	if err != nil {
		return
	}

	//获取当前提交的文件树
	tree, err := commit.Tree()
	if err != nil {
		return
	}

	// 获取指定Path的文件树
	if path != "" {
		tree, err = tree.Tree(path)
		if err != nil {
			return
		}
	}

	//组装Files
	res := make([]*File, len(tree.Entries))
	for i, v := range tree.Entries {
		res[i] = &File{
			Name:   v.Name,
			IsFile: v.Mode.IsFile(),
			Path:   path + v.Name,
		}
		if path != "" {
			res[i].Path = path + "/" + v.Name
		}
		if v.Mode.IsFile() {
			file, _ := tree.File(v.Name)
			res[i].Size = file.Size
		}
	}
	return res, nil
}
