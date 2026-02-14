package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/zichouu/go-upgrade/pkg/color"
	"github.com/zichouu/go-upgrade/pkg/exist"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func upGrade(path string) error {
	fmt.Println(color.Purple, "尝试", path, color.Reset)
	// git pull
	joinGit := filepath.Join(path, ".git")
	if exist.Bool(joinGit) {
		fmt.Println(color.Blue, "执行", path, "git pull", color.Reset)
		runCmd(path, "git", "pull")
	}
	// pnpm i
	joinPnpm := filepath.Join(path, "pnpm-lock.yaml")
	if exist.Bool(joinPnpm) {
		fmt.Println(color.Blue, "执行", path, "pnpm i", color.Reset)
		g.Go(func() error {
			err := runCmd(path, "pnpm", "i")
			return err
		})
		fmt.Println(color.Blue, "执行", path, "pnpm outdated", color.Reset)
		g.Go(func() error {
			err := runCmd(path, "pnpm", "outdated")
			return err
		})
		g.Wait()
	}
	return nil
}

func runCmd(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	fmt.Println(color.Green, "完成", dir, name, strings.Join(args, " "), color.Reset)
	if err != nil {
		fmt.Println(color.Red, string(out), color.Reset)
	} else {
		fmt.Println(string(out))
	}
	return nil
}
