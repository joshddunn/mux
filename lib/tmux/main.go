package tmux

import (
	"fmt"
	"strings"
)

type Tmux struct {
	Commands []string `exhaustruct:"optional"`
}

func (t *Tmux) append(command string) {
	t.Commands = append(t.Commands, command)
}

func (t *Tmux) GetCommand() string {
	commandString := strings.Join(t.Commands[:], " \\; ")
	return fmt.Sprintf("tmux %s \\;", commandString)
}

func (t *Tmux) ListSessions() {
	t.append("ls")
}

func (t *Tmux) AttachSession(target string) {
	t.append(fmt.Sprintf("attach-session -t %s", target))
}

func (t *Tmux) KillSession(target string) {
	t.append(fmt.Sprintf("kill-session -t %s", target))
}

func (t *Tmux) NewSession(target string) {
	t.append(fmt.Sprintf("new-session -t %s", target))
}

func (t *Tmux) KillWindow(target int) {
	t.append(fmt.Sprintf("kill-window -t %d", target))
}

func (t *Tmux) MoveWindow(target int, destination int) {
	t.append(fmt.Sprintf("move-window -s %d -t %d", target, destination))
}

func (t *Tmux) NewWindow(target int) {
	t.append(fmt.Sprintf("new-window -t %d", target))
}

func (t *Tmux) SelectWindow(target int) {
	t.append(fmt.Sprintf("select-window -t %d", target))
}

func (t *Tmux) KillPane(target int) {
	t.append(fmt.Sprintf("kill-pane -t %d", target))
}

func (t *Tmux) SelectPane(target int) {
	t.append(fmt.Sprintf("select-pane -t %d", target))
}

func (t *Tmux) SendKeys(keys string) {
	t.append(fmt.Sprintf("send-keys '%s'", keys))
}

func (t *Tmux) SendKeysEnter(keys string) {
	t.append(fmt.Sprintf("send-keys '%s' Enter", keys))
}
