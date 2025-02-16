package builder

import (
	"log"
	"mux/lib/config"
	"mux/lib/tmux"
)

func StartSession(name string) {
	t := tmux.Initialize(name)
	t.HasSession()

	config := config.Get()

	if t.Run() == nil {
		t.AttachSession()
		t.Exec()
	} else {
		session := findSession(name, config.Sessions)
		if session == nil {
			log.Fatal("Session doesn't exist in configuration")
		}

		buildSession(*session)
	}
}

func StopSession(name string) {
	t := tmux.Initialize(name)
	t.KillSession()
	if t.Run() != nil {
		log.Fatal("Session doesn't exist")
	}
}
