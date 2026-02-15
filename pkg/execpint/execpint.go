package execpint

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/zichouu/go-upgrade/pkg/color"
	"github.com/zichouu/go-upgrade/pkg/exist"
)

func Run(dir string, command string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}
	cmd.Dir = dir
	errColor := color.BGGreen
	fmt.Println(color.BGBlue, "执行", command, color.Reset)
	out, err := cmd.CombinedOutput()
	if err != nil {
		errColor = color.BGRed
	}
	fmt.Println(errColor, "完成", dir, command, color.Reset)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
	return err
}

func File(path string, filename string, command string) error {
	join := filepath.Join(path, filename)
	if exist.Bool(join) {
		err := Run(path, command)
		return err
	}
	return nil
}
