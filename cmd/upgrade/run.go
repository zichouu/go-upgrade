package main

import (
	"fmt"

	"github.com/zichouu/go-pkg/color"
	"github.com/zichouu/go-pkg/exe"
)

func run(path string) error {
	fmt.Println(color.BgPurple, "尝试", path, color.Reset)
	if PathGit {
		// git pull
		exe.IfExist(path, ".git", "git pull")
	}
	if PathPnpm {
		// pnpm i
		exe.IfExist(path, "pnpm-lock.yaml", "pnpm i")
		// pnpm outdated
		err := exe.IfExist(path, "pnpm-lock.yaml", "pnpm outdated")
		if err != nil {
			errPath = append(errPath, path)
		}
	}
	return nil
}
