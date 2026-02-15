package main

import (
	"github.com/charmbracelet/huh"
	"github.com/zichouu/go-upgrade/pkg/execpint"
	"golang.org/x/sync/errgroup"
)

var errPath []string

func up() {
	var g errgroup.Group
	var upPath []string
	upErr := true
	if len(errPath) > 0 {
		err := huh.NewMultiSelect[string]().
			Title("pnpm up --latest ?").
			Options(huh.NewOptions(errPath...)...).
			Value(&upPath).
			Run()
		if err != nil {
			upErr = false
		}
	}
	if len(upPath) > 0 && upErr {
		for _, value := range upPath {
			g.Go(func() error {
				execpint.Run(value, "pnpm up --latest")
				return nil
			})
		}
	}
	g.Wait()
}
