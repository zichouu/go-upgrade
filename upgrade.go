package main

import (
	"fmt"
	"path/filepath"

	"github.com/zichouu/go-upgrade/pkg/color"
	"github.com/zichouu/go-upgrade/pkg/execpint"
	"github.com/zichouu/go-upgrade/pkg/exist"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func upGrade(path string) error {
	fmt.Println(color.BGPurple, "尝试", path, color.Reset)
	// git pull
	joinGit := filepath.Join(path, ".git")
	if exist.Bool(joinGit) {
		fmt.Println(color.BGBlue, "执行", path, "git pull", color.Reset)
		execpint.Run(path, "git", "pull")
	}
	// pnpm i
	joinPnpm := filepath.Join(path, "pnpm-lock.yaml")
	if exist.Bool(joinPnpm) {
		fmt.Println(color.BGBlue, "执行", path, "pnpm i", color.Reset)
		g.Go(func() error {
			err := execpint.Run(path, "pnpm", "i")
			return err
		})
		fmt.Println(color.BGBlue, "执行", path, "pnpm outdated", color.Reset)
		g.Go(func() error {
			err := execpint.Run(path, "pnpm", "outdated")
			return err
		})
		g.Wait()
	}
	return nil
}
