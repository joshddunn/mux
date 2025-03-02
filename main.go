package main

import (
	"flag"
	"fmt"
	"log"
	"mux/embed"
	"mux/lib/builder"
	"mux/lib/config"
	"os"
)

const (
	Completion = "completion"
	Config     = "config"
	List       = "list"
	Start      = "start"
	Stop       = "stop"
)

var (
	version = "dev"
)

func main() {
	log.SetFlags(0)
	args := os.Args[1:]

	versionFlag := flag.Bool("version", false, "mux version")

	flag.Parse()

	if *versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	if len(args) == 0 {
		log.Fatal("Please enter a command.")
	}

	switch args[0] {
	case Completion:
		completion(args[1])
	case Config:
		config.EditConfig()
	case List:
		listSessions()
	case Start:
		startSession(args[1:])
	case Stop:
		stopSession(args[1:])
	default:
		log.Fatal("Invalid command.")
	}

	os.Exit(0)
}

func startSession(args []string) {
	if len(args) != 1 {
		log.Fatal("Invalid command. Usage: mux start <session>")
	}

	builder.StartSession(args[0])
}

func stopSession(args []string) {
	if len(args) != 1 {
		log.Fatal("Invalid command. Usage: mux stop <session>")
	}

	builder.StopSession(args[0])
}

func listSessions() {
	sessions := config.Get().Sessions

	for _, s := range sessions {
		fmt.Println(s.Name)
	}
}

func completion(shell string) {
	switch shell {
	case "zsh":
		fmt.Println(embed.MuxZsh)
	default:
		message := fmt.Sprintf("Completion is not available for `%s` shell.", shell)
		log.Fatal(message)
	}
}
