package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"

	"golang.org/x/sync/errgroup"
)

func main() {
	check()
	root := "."
	if len(os.Args) >= 2 {
		root = os.Args[1]
	}
	var g errgroup.Group
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
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
	g.Wait()
	// pnpm up --latest
	up()
}
