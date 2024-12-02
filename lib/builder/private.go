package builder

import (
	"math"
	"mux/lib/config"
	"mux/lib/tmux"
)

func findSession(name string, sessions []config.Session) *config.Session {
	var session *config.Session
	for _, s := range sessions {
		if s.Name == name {
			session = &s
			break
		}
	}

	return session
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
				buildBarsLayout(t, window, i, tmux.Horizontal)
			case config.Columns:
				buildBarsLayout(t, window, i, tmux.Vertical)
			}

			if i == 0 {
				t.KillPane(1)
			}

			if pane.Command != "" {
				t.SendKeys(pane.Command, *pane.Execute)
			}
		}

		t.SelectPane(1)
	}

	t.KillWindow(1)

	baseIndex := 1
	if *session.ZeroIndex {
		baseIndex = 0
	}

	for i := range session.Windows {
		t.MoveWindow(i+2, i+baseIndex)
	}

	t.SelectWindow(*session.SelectWindow)
	t.Exec()
}

func buildBarsLayout(t *tmux.Tmux, window config.Window, paneIndex int, split tmux.Split) {
	pane := window.Panes[paneIndex]
	percent := splitPercent(paneIndex, len(window.Panes))
	t.SplitWindow(pane.Dir, split, percent)
}

func buildDefaultLayout(t *tmux.Tmux, window config.Window, paneIndex int) {
	pane := window.Panes[paneIndex]
	split := tmux.Horizontal
	percent := *window.SplitPercent

	if paneIndex > 1 {
		split = tmux.Vertical
		percent = splitPercent(paneIndex, len(window.Panes))
	}

	t.SplitWindow(pane.Dir, split, percent)
}

func splitPercent(index int, count int) int {
	return int(math.Ceil(100 - 100/float64(count-index+1)))
}
