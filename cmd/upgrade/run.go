package main

import (
	"fmt"

	"github.com/zichouu/go-upgrade/pkg/color"
	"github.com/zichouu/go-upgrade/pkg/execpint"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func run(path string) error {
	fmt.Println(color.BGPurple, "尝试", path, color.Reset)
	// git pull
	execpint.File(path, ".git", "git", "pull")
	// pnpm i
	g.Go(func() error {
		execpint.File(path, "pnpm-lock.yaml", "pnpm", "i")
		return nil
	})
	// pnpm outdated
	g.Go(func() error {
		err := execpint.File(path, "pnpm-lock.yaml", "pnpm", "outdated")
		if err != nil {
			errPath = append(errPath, path)
		}
		return nil
	})
	g.Wait()
	return nil
}
