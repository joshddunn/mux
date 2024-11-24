package main

// https://github.com/alexflint/go-arg
// https://github.com/alecthomas/kong
// https://github.com/spf13/cobra
// https://github.com/urfave/cli
// https://github.com/spf13/viper

import (
	"mux/lib/tmux"
)

func main() {
	tmux := tmux.Initialize("bob")
	tmux.AttachSessionIfExists()

	tmux.NewSession()
	tmux.NewWindow("bob2")
	tmux.ExecCommand()
}
