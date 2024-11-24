package tmux

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

type Tmux struct {
	SessionName string
	binary      string
	commands    []string
}

func Initialize(sessionName string) *Tmux {
	binary, lookErr := exec.LookPath("tmux")
	if lookErr != nil {
		panic(lookErr)
	}

	return &Tmux{
		SessionName: sessionName,
		binary:      binary,
		commands:    []string{},
	}
}

// private

func (t *Tmux) append(cmds ...string) {
	t.commands = append(t.commands, append(cmds, ";")...)
}

func (t *Tmux) getCommand() []string {
	return append([]string{"tmux"}, t.commands...)
}

// public

func (t *Tmux) AttachIfSessionExists() {
	cmd := exec.Command("tmux", "has-session", "-t", t.SessionName)
	err := cmd.Run()

	if err == nil {
		t.AttachSession()
		t.ExecCommand()
	}
}

func (t *Tmux) ExecCommand() {
	log.Fatal(syscall.Exec(t.binary, t.getCommand(), os.Environ()))
}

func (t *Tmux) ListSessions() {
	t.append("ls")
}

func (t *Tmux) AttachSession() {
	t.append("attach-session", "-t", t.SessionName)
}

func (t *Tmux) KillSession() {
	t.append("kill-session", "-t", t.SessionName)
}

func (t *Tmux) NewSession() {
	t.append("new-session", "-s", t.SessionName)
}

func (t *Tmux) KillWindow(target int) {
	t.append("kill-window", "-t", strconv.Itoa(target))
}

func (t *Tmux) MoveWindow(target int, destination int) {
	t.append("move-window", "-s", strconv.Itoa(target), "-t", strconv.Itoa(destination))
}

func (t *Tmux) NewWindow(name string) {
	t.append("new-window", "-n", name)
}

func (t *Tmux) SelectWindow(target int) {
	t.append("select-window", "-t", strconv.Itoa(target))
}

func (t *Tmux) KillPane(target int) {
	t.append("kill-pane", "-t", strconv.Itoa(target))
}

func (t *Tmux) SelectPane(target int) {
	t.append("select-pane", "-t", strconv.Itoa(target))
}

func (t *Tmux) SendKeys(keys string) {
	t.append("send-keys", keys)
}

func (t *Tmux) SendKeysEnter(keys string) {
	t.append("send-keys", keys, "Enter")
}
