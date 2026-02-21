package main

import (
	"github.com/charmbracelet/huh"
	"github.com/zichouu/go-pkg/exe"
	"golang.org/x/sync/errgroup"
)

var errPath []string

func up() {
	var g errgroup.Group
	if len(errPath) > 0 {
		var upPath []string
		err := huh.NewMultiSelect[string]().
			Title("pnpm up --latest ?").
			Options(huh.NewOptions(errPath...)...).
			Value(&upPath).
			Run()
		if len(upPath) > 0 && err == nil {
			for _, v := range upPath {
				g.Go(func() error {
					_, _ = exe.Run(v, []string{}, "pnpm", "up", "--latest")
					return nil
				})
			}
		}
	}
	_ = g.Wait()
}
