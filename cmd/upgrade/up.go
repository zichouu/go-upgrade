package main

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/zichouu/go-upgrade/pkg/execpint"
	"golang.org/x/sync/errgroup"
)

var errPath []string

func up() {
	var g errgroup.Group
	up := false
	if len(errPath) > 0 {
		fmt.Println(errPath)
		huh.NewConfirm().
			Title("pnpm up --latest ?").
			Value(&up).
			Run()
	}
	if up {
		for _, value := range errPath {
			g.Go(func() error {
				execpint.Run(value, "pnpm up --latest")
				return nil
			})
		}
	}
	g.Wait()
}
