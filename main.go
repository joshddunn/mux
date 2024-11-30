package main

// https://github.com/alexflint/go-arg
// https://github.com/alecthomas/kong
// https://github.com/spf13/cobra
// https://github.com/urfave/cli
// https://github.com/spf13/viper

import (
	"log"
	"mux/lib/config"
	"mux/lib/tmux"
	"os"
)

const (
	Start = "start"
	Stop  = "stop"
	Help  = "help"
)

func main() {
	log.SetFlags(0)

	args := os.Args[1:]

	RunCommand(args)
}

func RunCommand(args []string) {
	if len(args) != 2 {
		log.Fatal("Invalid command line args")
	}

	switch args[0] {
	case Start:
		StartCommand(args[1])
	case Stop:
		StopCommand(args[1])
	default:
		log.Fatal("Invalid command line argument")
	}
}

func StartCommand(name string) {
	t := tmux.Initialize(name)
	t.HasSession()
	err := t.Run()

	if err == nil {
		t.AttachSession()
		err = t.Exec()
		if err != nil {
			log.Fatal("Something went wrong")
		}
	} else {
		c, err := config.Get()

		if err != nil {
			log.Fatal(err)
		}

		var session *config.Session
		for _, s := range c.Sessions {
			if s.Name == name {
				session = &s
				break
			}
		}

		if session == nil {
			log.Fatal("Session doesn't exist in configuration")
		}

		BuildSession(*session)
	}
}

func StopCommand(name string) {
	t := tmux.Initialize(name)
	t.HasSession()
	t.KillSession()
	err := t.Run()
	if err != nil {
		log.Fatal("Session doesn't exist")
	}
}

func BuildSession(session config.Session) {
	log.Print(session)
}
