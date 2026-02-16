package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/zichouu/go-pkg/color"
)

var PathGit = true
var PathPnpm = true

func check() {
	if _, err := exec.LookPath("git"); err != nil {
		PathGit = false
		fmt.Println(color.BgRed, err, color.Reset)
	}
	if _, err := exec.LookPath("pnpm"); err != nil {
		PathPnpm = false
		fmt.Println(color.BgRed, err, color.Reset)
	}
	if !PathGit && !PathPnpm {
		os.Exit(1)
	}
}
