package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/zichouu/go-pkg/color"
)

var CanUseList = []string{}

func check() {
	checkList := []string{"git", "pnpm"}
	for _, v := range checkList {
		_, err := exec.LookPath(v)
		if err != nil {
			fmt.Println(color.BgRed, err, color.Reset)
		} else {
			CanUseList = append(CanUseList, v)
		}
	}
	if len(CanUseList) == 0 {
		os.Exit(1)
	}
}
