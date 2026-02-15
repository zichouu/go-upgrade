package main

import (
	"fmt"

	"github.com/zichouu/go-upgrade/pkg/color"
	"github.com/zichouu/go-upgrade/pkg/exe"
)

func run(path string) error {
	fmt.Println(color.BGPurple, "尝试", path, color.Reset)
	if PathGit {
		// git pull
		exe.File(path, ".git", "git pull")
	}
	if PathPnpm {
		// pnpm i
		exe.File(path, "pnpm-lock.yaml", "pnpm i")
		// pnpm outdated
		err := exe.File(path, "pnpm-lock.yaml", "pnpm outdated")
		if err != nil {
			errPath = append(errPath, path)
		}
	}
	return nil
}
