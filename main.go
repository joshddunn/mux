package main

// https://github.com/alexflint/go-arg
// https://github.com/alecthomas/kong
// https://github.com/spf13/cobra
// https://github.com/urfave/cli
// https://github.com/spf13/viper

import (
	"fmt"
	"mux/lib/tmux"
)

func main() {
	// binary, lookErr := exec.LookPath("tmux")
	// if lookErr != nil {
	// 	panic(lookErr)
	// }

	// args := []string{"tmux", "a", "-t", "dev"}

	tmux := &tmux.Tmux{}

	tmux.NewSession("dev")
	tmux.MoveWindow(1, 2)

	cmd := tmux.GetCommand()

	fmt.Print(cmd)

	// err := syscall.Exec(binary, args, os.Environ())

	// if err != nil {
	// 	log.Fatal(err)
	// }
}
