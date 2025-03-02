package builder

import (
	"fmt"
	"log"
	"mux/lib/config"
	"mux/lib/tmux"
)

func StartSession(name string) {
	t := tmux.Initialize(name)
	t.HasSession()

	c := config.Get()

	if t.Run() == nil {
		t.AttachSession()
		t.Exec()
	} else {
		session := findSession(name, c.Sessions)
		if session == nil {
			message := fmt.Sprintf("Session `%s` not defined in ~/%s", name, config.File)
			log.Fatal(message)
		}

		buildSession(*session)
	}
}

func StopSession(name string) {
	t := tmux.Initialize(name)
	t.KillSession()
	if t.Run() != nil {
		message := fmt.Sprintf("`%s` is not an active session", name)
		log.Fatal(message)
	}
}
