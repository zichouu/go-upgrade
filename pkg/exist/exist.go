package exist

import "os"

func Bool(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
