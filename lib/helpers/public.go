package helpers

import (
	"os"
	"strings"
)

func Pointer[T any](d T) *T {
	return &d
}

func DirectoryFullpath(dir string) string {
	return strings.Replace(dir, "~", HomeDir(), 1)
}

func DirectoryExists(dir string) bool {
	fullpath := DirectoryFullpath(dir)
	stat, err := os.Stat(fullpath)

	if err == nil {
		return stat.IsDir()
	}

	if os.IsNotExist(err) {
		return false
	}

	panic(err)
}

func HomeDir() string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return homedir
}
