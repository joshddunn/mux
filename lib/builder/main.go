package builder

import (
	"log"
	"math"
	"mux/lib/config"
	"mux/lib/tmux"
)

func StartSession(name string) {
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

		buildSession(*session)
	}
}

func StopSession(name string) {
	t := tmux.Initialize(name)
	t.KillSession()
	err := t.Run()
	if err != nil {
		log.Fatal("Session doesn't exist")
	}
}

func buildSession(session config.Session) {
	t := tmux.Initialize(session.Name)
	t.NewSession()

	for _, window := range session.Windows {
		t.NewWindow(window.Name)

		for i, pane := range window.Panes {

			switch *window.Layout {
			case config.Default:
				buildDefaultLayout(t, window, i)
			case config.Rows:
				buildBarsLayout(t, tmux.Horizontal, window.Panes, i)
			case config.Columns:
				buildBarsLayout(t, tmux.Vertical, window.Panes, i)
			}

			if pane.Command != "" {
				t.SendKeys(pane.Command, *pane.Execute)
			}
		}

		t.SelectPane(1)
	}

	t.KillWindow(1)

	for i := range session.Windows {
		if *session.ZeroIndex {
			t.MoveWindow(i+2, i)
		} else {
			t.MoveWindow(i+2, i+1)
		}
	}

	t.SelectWindow(*session.SelectWindow)

	err := t.Exec()
	if err != nil {
		log.Fatal("Something went wrong")
	}
}

func buildBarsLayout(t *tmux.Tmux, split tmux.Split, panes []config.Pane, paneIndex int) {
	pane := panes[paneIndex]
	percent := splitPercent(paneIndex, len(panes))
	t.SplitWindow(pane.Dir, split, percent)

	if paneIndex == 0 {
		t.KillPane(1)
	}
}

func buildDefaultLayout(t *tmux.Tmux, window config.Window, paneIndex int) {
	pane := window.Panes[paneIndex]

	switch paneIndex {
	case 0:
		t.SplitWindow(pane.Dir, tmux.Vertical, *window.SplitPercent)
		t.KillPane(1)
	case 1:
		t.SplitWindow(pane.Dir, tmux.Horizontal, *window.SplitPercent)
	default:
		percent := splitPercent(paneIndex, len(window.Panes))
		t.SplitWindow(pane.Dir, tmux.Vertical, percent)
	}
}

func splitPercent(index int, count int) int {
	return int(math.Ceil(100 - 100/float64(count-index+1)))
}
