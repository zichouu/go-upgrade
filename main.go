package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	filepath.WalkDir("../test", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			// 跳过某些目录
			if d.Name() == ".git" || d.Name() == "node_modules" {
				return filepath.SkipDir
			}
			upGrade(path)
		}
		return nil
	})
}

func upGrade(path string) {
	fmt.Println(colorPurple, path, "尝试目录", colorReset)
	// git pull
	joinGit := filepath.Join(path, ".git")
	if isExist(joinGit) {
		fmt.Println(colorBlue, path, "git pull 执行", colorReset)
		runCmd(path, "git", "pull")
	}
	// pnpm i
	joinPnpm := filepath.Join(path, "pnpm-lock.yaml")
	if isExist(joinPnpm) {
		fmt.Println(colorBlue, path, "pnpm i 执行", colorReset)
		runCmd(path, "pnpm", "i")
	}
}

func runCmd(dir, name, args string) {
	cmd := exec.Command(name, args)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(colorGreen, dir, name, args, "完成", colorReset)
	fmt.Println(string(out))
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

var colorReset = "\033[0m"
var colorGreen = "\033[30;42m"
var colorBlue = "\033[97;44m"
var colorPurple = "\033[97;45m"
