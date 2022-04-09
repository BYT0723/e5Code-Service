package gitx

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"time"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

type File struct {
	Name       string
	IsFile     bool
	Mode       string
	Size       int64 //unit:Byte
	Path       string
	CommitHash string
	CommitInfo *CommitInfo
	Children   []*File
}

type CommitInfo struct {
	Message string
	Author  string
	When    time.Time
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
func ListFile(rep *git.Repository, path string, isRecures, isWork bool) ([]*File, error) {
	if !isWork {
		// 获取Head指针
		ref, err := rep.Head()
		if err != nil {
			return nil, err
		}

		// 获取Head的Hash的提交
		commit, err := rep.CommitObject(ref.Hash())
		if err != nil {
			return nil, err
		}

		//获取当前提交的文件树
		tree, err := commit.Tree()
		if err != nil {
			return nil, err
		}

		// 获取指定Path的文件树
		if path != "" {
			tree, err = tree.Tree(path)
			if err != nil {
				return nil, err
			}
		}

		res := listCommitTree(rep, tree, path, isRecures)
		return res, nil
	} else {
		w, err := rep.Worktree()
		if err != nil {
			return nil, err
		}
		res := listWorkTree(w, path, isRecures)
		return res, nil
	}
}

func listCommitTree(rep *git.Repository, tree *object.Tree, path string, isRecures bool) []*File {
	res := make([]*File, len(tree.Entries))
	for i, v := range tree.Entries {
		// 初始化File
		res[i] = &File{
			Name:   v.Name,
			Mode:   "Dir",
			IsFile: v.Mode.IsFile(),
			Path:   path + v.Name,
		}

		// 判断是否为根目录
		if path != "" {
			res[i].Path = path + "/" + v.Name
		}

		// 判断是否为文件
		if v.Mode.IsFile() {
			// 获取文件size
			file, _ := tree.File(v.Name)

			// 获取文件Log（获取最新commit）
			commitIter, _ := rep.Log(&git.LogOptions{
				FileName: &res[i].Path,
			})
			commit, _ := commitIter.Next()

			res[i].Size = file.Size
			res[i].Mode = "File"
			res[i].CommitHash = commit.Hash.String()
			res[i].CommitInfo = &CommitInfo{
				Message: strings.TrimSuffix(commit.Message, "\n"),
				Author:  commit.Author.Name,
				When:    commit.Author.When.UTC(),
			}
		} else if isRecures {
			subTree, _ := tree.Tree(v.Name)
			res[i].Children = listCommitTree(rep, subTree, res[i].Path, isRecures)
		}
	}
	for i := 0; i < len(res); i++ {
		if !res[i].IsFile {
			res[i].Size = CountSize(res[i])
			GetCommit(res[i])
		}
	}
	return res
}

func listWorkTree(tree *git.Worktree, path string, isRecures bool) []*File {
	files, _ := tree.Filesystem.ReadDir(path)
	res := make([]*File, len(files))
	if path != "" {
		path += "/"
	}
	for i, f := range files {
		res[i] = &File{
			Name:     f.Name(),
			Size:     f.Size(),
			Mode:     f.Mode().String(),
			Path:     path + f.Name(),
			IsFile:   !f.IsDir(),
			Children: []*File{},
		}
		if f.IsDir() && isRecures {
			res[i].Children = listWorkTree(tree, res[i].Path, isRecures)
		}
	}
	return res
}

// 计算文件夹的总大小
func CountSize(f *File) int64 {
	var sum int64 = 0
	for _, v := range f.Children {
		if v.IsFile {
			sum += v.Size
		} else {
			sum += CountSize(v)
		}
	}
	return sum
}

// 获取文件夹下最新的一条提交记录
func GetCommit(f *File) {
	when, _ := time.Parse("2006-01-02", "1970-01-01")
	for _, v := range f.Children {
		if v.IsFile {
			if v.CommitInfo.When.After(when) {
				f.CommitHash = v.CommitHash
				f.CommitInfo = v.CommitInfo
			}
		} else {
			GetCommit(v)
		}
	}
}

func ReadFile(rep *git.Repository, path string, isWork bool) (string, error) {
	if !isWork {
		// 获取Head指针
		ref, err := rep.Head()
		if err != nil {
			return "", err
		}

		// 获取Head的Hash的提交
		commit, err := rep.CommitObject(ref.Hash())
		if err != nil {
			return "", err
		}

		//获取当前提交的文件树
		tree, err := commit.Tree()
		if err != nil {
			return "", err
		}
		file, err := tree.File(path)
		if err != nil {
			return "", err
		}
		if file.Size > 10*1024*1024 {
			return "", errors.New("FileSizeOverSize")
		}
		return file.Contents()

	} else {
		//获取工作树
		w, err := rep.Worktree()
		if err != nil {
			return "", err
		}
		file, err := w.Filesystem.OpenFile(path, os.O_RDONLY, 0755)
		if err != nil {
			return "", err
		}
		body, err := ioutil.ReadAll(file)
		return string(body), err
	}
}

func CreateFile(rep *git.Repository, path string) (err error) {
	w, err := rep.Worktree()
	if err != nil {
		return
	}
	_, err = w.Filesystem.Create(path)
	return
}

func UpdateFile(rep *git.Repository, path string, body []byte) (err error) {
	//获取工作树
	w, err := rep.Worktree()
	if err != nil {
		return
	}

	opt := os.O_WRONLY | os.O_TRUNC
	// 打开指定文件
	file, err := w.Filesystem.OpenFile(path, opt, 0640)
	if err != nil {
		return
	}
	if _, err = file.Write(body); err != nil {
		return
	}
	if _, err = w.Add(path); err != nil {
		return
	}
	return
}

// 删除指定文件或文件夹
func DeleteFile(rep *git.Repository, path string) (err error) {
	w, err := rep.Worktree()
	if err != nil {
		return
	}
	_, err = w.Remove(path)
	return
}

// mv
func MoveFile(rep *git.Repository, oldpath, newpath string) (err error) {
	w, err := rep.Worktree()
	if err != nil {
		return
	}
	return w.Filesystem.Rename(oldpath, newpath)
}

func MkDir(rep *git.Repository, path string) (err error) {
	w, err := rep.Worktree()
	if err != nil {
		return
	}
	return w.Filesystem.MkdirAll(path, os.ModeDir)
}

type FileStatus struct {
	Path   string
	Status string
}

func GitStatus(rep *git.Repository) (res []*FileStatus, err error) {
	w, err := rep.Worktree()
	if err != nil {
		return
	}
	status, err := w.Status()
	if err != nil {
		return
	}
	for k, v := range status {
		res = append(res, &FileStatus{
			Path:   k,
			Status: string(v.Staging),
		})
	}
	return
}

type CommitOption struct {
	FilePaths []string
	Msg       string
	Author    string
	Email     string
	Remote    string
	*http.BasicAuth
}

func Commit(rep *git.Repository, opt *CommitOption) (err error) {
	//获取工作树
	w, err := rep.Worktree()
	if err != nil {
		return
	}
	// add change
	for _, path := range opt.FilePaths {
		w.Add(path)
	}

	// 提交更改
	if _, err = w.Commit(opt.Msg, &git.CommitOptions{
		Author: &object.Signature{
			Name:  opt.Author,
			Email: opt.Email,
			When:  time.Now(),
		},
	}); err != nil {
		return
	}
	if opt.Remote == "" {
		opt.Remote = "origin"
	}
	if err = rep.Push(&git.PushOptions{
		RemoteName: opt.Remote,
		Auth:       opt.BasicAuth,
	}); err != nil {
		return
	}
	return
}
