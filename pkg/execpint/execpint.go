package execpint

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/zichouu/go-upgrade/pkg/color"
	"github.com/zichouu/go-upgrade/pkg/exist"
)

func Run(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	errColor := color.BGGreen
	fmt.Println(color.BGBlue, "执行", dir, name, strings.Join(args, " "), color.Reset)
	out, err := cmd.CombinedOutput()
	if err != nil {
		errColor = color.BGRed
	}
	fmt.Println(errColor, "完成", dir, name, strings.Join(args, " "), color.Reset)
	fmt.Println(string(out))
	return err
}

func File(path string, filename string, name string, args ...string) error {
	join := filepath.Join(path, filename)
	if exist.Bool(join) {
		err := Run(path, name, args...)
		return err
	}
	return nil
}
