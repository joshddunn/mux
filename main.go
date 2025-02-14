package main

import (
	"flag"
	"log"
	"mux/lib/builder"
	"mux/lib/config"
	"os"
)

const (
	Config  = "config"
	Help    = "help"
	Start   = "start"
	Stop    = "stop"
	Version = "v0.1.2"
)

func main() {
	log.SetFlags(0)
	cmd := os.Args[1]
	args := os.Args[2:]

	versionPtr := flag.Bool("version", false, "mux version")

	flag.Parse()

	if *versionPtr {
		log.Print(Version)
		os.Exit(0)
	}

	switch cmd {
	case Config:
		config.EditConfig()
	case Help:
		log.Print("https://github.com/joshddunn/mux")
	case Start:
		startSession(args)
	case Stop:
		stopSession(args)
	default:
		log.Fatal("Invalid command line argument")
	}
}

func startSession(args []string) {
	if len(args) != 1 {
		log.Fatal("Invalid command line args")
	}

	builder.StartSession(args[0])
}

func stopSession(args []string) {
	if len(args) != 1 {
		log.Fatal("Invalid command line args")
	}

	builder.StopSession(args[0])
}
