package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"golang.org/x/sync/errgroup"
)

func main() {
	check()
	var g errgroup.Group
	args := "."
	if len(os.Args) >= 2 {
		args = os.Args[1]
	}
	filepath.WalkDir(args, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			// run() 前跳过某些目录 确定是无效目录
			if d.Name() == ".git" || d.Name() == "node_modules" {
				return filepath.SkipDir
			}
			g.Go(func() error {
				err := run(path)
				return err
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
	if err := g.Wait(); err != nil {
		fmt.Println("错误", err)
	}
	// pnpm up --latest
	if PathPnpm {
		up()
	}
}
