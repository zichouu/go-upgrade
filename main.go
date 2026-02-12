package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group
	filepath.WalkDir("../test", func(path string, d fs.DirEntry, err error) error {
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
		fmt.Println(err)
	}
}

func upGrade(path string) error {
	fmt.Println(colorPurple, path, "尝试目录", colorReset)
	// git pull
	joinGit := filepath.Join(path, ".git")
	if isExist(joinGit) {
		fmt.Println(colorBlue, path, "git pull 执行", colorReset)
		err := runCmd(path, "git", "pull")
		if err != nil {
			return err
		}
	}
	// pnpm i
	joinPnpm := filepath.Join(path, "pnpm-lock.yaml")
	if isExist(joinPnpm) {
		fmt.Println(colorBlue, path, "pnpm i 执行", colorReset)
		err := runCmd(path, "pnpm", "i")
		if err != nil {
			return err
		}
	}
	return nil
}

func runCmd(dir, name, args string) error {
	cmd := exec.Command(name, args)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(colorGreen, dir, name, args, "完成", colorReset)
	fmt.Println(string(out))
	return nil
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

var colorReset = "\033[0m"
var colorGreen = "\033[30;42m"
var colorBlue = "\033[97;44m"
var colorPurple = "\033[97;45m"
