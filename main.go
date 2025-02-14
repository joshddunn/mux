package main

import (
	"log"
	"mux/lib/builder"
	"mux/lib/config"
	"os"
)

const (
	Config = "config"
	Start  = "start"
	Stop   = "stop"
)

func main() {
	log.SetFlags(0)
	args := os.Args[1:]

	switch args[0] {
	case Config:
		config.EditConfig()
	case Start:
		startSession(args)
	case Stop:
		stopSession(args)
	default:
		log.Fatal("Invalid command line argument")
	}
}

func startSession(args []string) {
	if len(args) != 2 {
		log.Fatal("Invalid command line args")
	}

	builder.StartSession(args[1])
}

func stopSession(args []string) {
	if len(args) != 2 {
		log.Fatal("Invalid command line args")
	}

	builder.StopSession(args[1])
}
