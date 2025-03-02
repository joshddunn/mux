#compdef mux
#definition manage tmux sessions

local -a subcmds

subcmds=(
  "config:Open config file (~/.mux.json) with $(echo $EDITOR)",
  'start:Start session in config file',
  'stop:Stop session in config file'
)

_arguments '--version[Current version]' \
  '*:: :->subcmds' && return 0

if (( CURRENT == 1 )); then
  _describe -t subcmds "mux subcommands" subcmds
  return
fi

case "$words[1]" in
  start)
    compadd $(mux list) ;;
  stop)
    compadd $(mux list) ;;
esac
