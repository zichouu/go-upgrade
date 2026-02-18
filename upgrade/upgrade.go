package main

import (
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"slices"

	"github.com/zichouu/go-pkg/check"
	"golang.org/x/sync/errgroup"
)

var CanUseList = []string{}

func main() {
	out, err := check.Path([]string{"git", "pnpm"})
	if err != nil {
		os.Exit(1)
	} else {
		CanUseList = out
	}
	root := "."
	if len(os.Args) >= 2 {
		root = os.Args[1]
	}
	var g errgroup.Group
	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			// run() 前跳过某些目录 确定是无效目录
			skipList := []string{".git", "node_modules"}
			if slices.Contains(skipList, d.Name()) {
				return filepath.SkipDir
			}
			g.Go(func() error {
				run(path)
				return nil
			})
			// run() 后 检测目录是否存在 .git 并跳过
			joinGit := filepath.Join(path, ".git")
			_, err := os.Stat(joinGit)
			if err == nil {
				return filepath.SkipDir
			}
		}
		return nil
	})
	if err != nil {
		slog.Error("WalkDir err", "err", err)
	}
	_ = g.Wait()
	// pnpm up --latest
	up()
}
