package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/zichouu/go-pkg/color"
	"github.com/zichouu/go-pkg/exe"
)

func run(path string) {
	fmt.Println(color.BgPurple, "尝试", path, color.Reset)
	if slices.Contains(CanUseList, "git") {
		// git pull
		_, err := exe.IfExist(path, ".git", []string{}, "git", "pull")
		if err == nil {
			if slices.Contains(CanUseList, "pnpm") {
				// pnpm i
				_, _ = exe.IfExist(path, "pnpm-lock.yaml", []string{}, "pnpm", "i")
				// pnpm outdated
				out, err := exe.IfExist(path, "pnpm-lock.yaml", []string{}, "pnpm", "outdated")
				outString := string(out)
				isOutDated := strings.Contains(outString, "Package") &&
					strings.Contains(outString, "Current") &&
					strings.Contains(outString, "Latest")
				if err != nil && isOutDated {
					errPath = append(errPath, path)
				}
			}
		}
	}
}
