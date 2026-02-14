package execpint

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/zichouu/go-upgrade/pkg/color"
)

func Run(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	errColor := color.BGGreen
	out, err := cmd.CombinedOutput()
	if err != nil {
		errColor = color.BGRed
	}
	fmt.Println(errColor, "完成", dir, name, strings.Join(args, " "), color.Reset)
	fmt.Println(string(out))
	return nil
}
