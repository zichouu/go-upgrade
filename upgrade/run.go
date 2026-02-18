package main

import (
	"fmt"
	"slices"

	"github.com/zichouu/go-pkg/color"
	"github.com/zichouu/go-pkg/exe"
)

func run(path string) {
	fmt.Println(color.BgPurple, "尝试", path, color.Reset)
	if slices.Contains(CanUseList, "git") {
		// git pull
		_ = exe.IfExist(path, ".git", []string{}, "git", "pull")
	}
	if slices.Contains(CanUseList, "pnpm") {
		// pnpm i
		_ = exe.IfExist(path, "pnpm-lock.yaml", []string{}, "pnpm", "i")
		// pnpm outdated
		err := exe.IfExist(path, "pnpm-lock.yaml", []string{}, "pnpm", "outdated")
		if err != nil {
			errPath = append(errPath, path)
		}
	}
}
