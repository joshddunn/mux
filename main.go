package main

// https://github.com/alexflint/go-arg
// https://github.com/alecthomas/kong
// https://github.com/spf13/cobra
// https://github.com/urfave/cli
// https://github.com/spf13/viper

import (
	"log"
	"mux/lib/tmux"
)

func main() {
	t := tmux.Initialize("bob")

	t.NewSession()
	t.NewWindow("bob2")
	t.SplitWindow("~/code", tmux.Horizontal, 50)
	log.Fatal(t.Exec())
}
