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
	Config     = "config"
	List       = "list"
	Start      = "start"
	Stop       = "stop"
	Completion = "completion"
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
		log.Print(version)
		os.Exit(0)
	}

	if len(args) == 0 {
		log.Fatal("Invalid command (no args)")
	}

	switch args[0] {
	case Config:
		config.EditConfig()
	case Start:
		startSession(args[1:])
	case Stop:
		stopSession(args[1:])
	case List:
		listSessions()
	case Completion:
		completion(args[1])
	default:
		log.Fatal(fmt.Sprintf("Invalid command (cant find %s)", args[0]))
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
		fmt.Print(embed.MuxZsh)
	default:
		fmt.Print("# noop")
	}
}
