package main

import (
	"log"
	"mux/lib/builder"
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

	if len(args) != 2 {
		log.Fatal("Invalid command line args")
	}

	switch args[0] {
	case Start:
		builder.StartSession(args[1])
	case Stop:
		builder.StopSession(args[1])
	default:
		log.Fatal("Invalid command line argument")
	}
}
