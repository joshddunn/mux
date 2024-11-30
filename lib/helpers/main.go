package helpers

import (
	"os"
	"strings"
)

func Pointer[T any](d T) *T {
	return &d
}

func DirectoryExists(dir string) bool {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	fullpath := strings.Replace(dir, "~", homedir, 1)
	stat, err := os.Stat(fullpath)

	if err == nil {
		return stat.IsDir()
	}

	if os.IsNotExist(err) {
		return false
	}

	panic(err)
}
