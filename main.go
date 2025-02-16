package main

import (
	_ "embed"
	"flag"
	"log"
	"mux/lib/builder"
	"mux/lib/config"
	"os"
)

const (
	Config = "config"
	Help   = "help"
	Start  = "start"
	Stop   = "stop"
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
		log.Fatal("Invalid command")
	}

	switch args[0] {
	case Config:
		config.EditConfig()
	case Help:
		log.Print("https://github.com/joshddunn/mux")
	case Start:
		startSession(args[1:])
	case Stop:
		stopSession(args[1:])
	default:
		log.Fatal("Invalid command")
	}
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
