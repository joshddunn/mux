package tmux

import (
	"fmt"
	"mux/lib/helpers"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func Initialize(sessionName string) *Tmux {
	binary, err := exec.LookPath("tmux")
	if err != nil {
		panic(err)
	}

	return &Tmux{
		SessionName: sessionName,
		binary:      binary,
		commands:    []string{},
	}
}

func (t *Tmux) Exec() {
	args := append([]string{"tmux"}, t.commands...)
	err := syscall.Exec(t.binary, args, os.Environ())
	if err != nil {
		panic(err)
	}
}

func (t *Tmux) Reset() {
	t.commands = []string{}
}

func (t *Tmux) Run() error {
	cmd := exec.Command("tmux", t.commands...)
	err := cmd.Run()

	t.Reset()

	return err
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

func (t *Tmux) HasSession() {
	t.append("has-session", "-t", t.SessionName)
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

func (t *Tmux) SplitWindow(dir string, split Split, percent int) {
	fullpath := helpers.DirectoryFullpath(dir)
	cmd := []string{"split-window", "-c", fullpath}

	switch split {
	case Horizontal:
		cmd = append(cmd, "-h")
	case Vertical:
		cmd = append(cmd, "-v")
	}

	if percent > 0 {
		cmd = append(cmd, "-l", fmt.Sprintf("%d%%", percent))
	}

	t.append(cmd...)
}

func (t *Tmux) KillPane(target int) {
	t.append("kill-pane", "-t", strconv.Itoa(target))
}

func (t *Tmux) SelectPane(target int) {
	t.append("select-pane", "-t", strconv.Itoa(target))
}

func (t *Tmux) SendKeys(keys string, enter bool) {
	args := []string{"send-keys", keys}

	if enter {
		args = append(args, "Enter")
	}

	t.append(args...)
}
