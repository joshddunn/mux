package tmux

type Split string

type Tmux struct {
	SessionName string
	binary      string
	commands    []string
}
