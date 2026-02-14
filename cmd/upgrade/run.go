package main

import (
	"fmt"

	"github.com/zichouu/go-upgrade/pkg/color"
	"github.com/zichouu/go-upgrade/pkg/execpint"
)

func run(path string) error {
	fmt.Println(color.BGPurple, "尝试", path, color.Reset)
	// git pull
	execpint.File(path, ".git", "git", "pull")
	// pnpm i
	execpint.File(path, "pnpm-lock.yaml", "pnpm", "i")
	// pnpm outdated
	err := execpint.File(path, "pnpm-lock.yaml", "pnpm", "outdated")
	if err != nil {
		errPath = append(errPath, path)
	}
	return nil
}
