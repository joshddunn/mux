package tmux

func (t *Tmux) append(cmds ...string) {
	t.commands = append(t.commands, append(cmds, ";")...)
}
