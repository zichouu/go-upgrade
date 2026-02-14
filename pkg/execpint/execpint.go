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
	out, err := cmd.CombinedOutput()
	fmt.Println(color.BGGreen, "完成", dir, name, strings.Join(args, " "), color.Reset)
	if err != nil {
		fmt.Println(color.Red, string(out), color.Reset)
	} else {
		fmt.Println(string(out))
	}
	return nil
}
