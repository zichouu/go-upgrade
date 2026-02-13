package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"golang.org/x/sync/errgroup"
)

func main() {
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
			// 跳过某些目录
			if d.Name() == ".git" || d.Name() == "node_modules" {
				return filepath.SkipDir
			}
			g.Go(func() error {
				err := upGrade(path)
				if err != nil {
					return err
				}
				return nil
			})
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		fmt.Println("错误", err)
	}
}
