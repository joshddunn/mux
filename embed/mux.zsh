#compdef mux
#autoload

_mux() {
  local -a subcmds

  subcmds=(
    "config:Open config file (~/.mux.json) with $(echo $EDITOR)",
    'list:List sessions in config file',
    'start:Start session in config file',
    'stop:Stop session in config file'
  )

  _arguments '--version[Current version]' \
    '*:: :->subcmds' && return 0

  if (( CURRENT == 1 )); then
    _describe "mux" subcmds
    return
  fi

  case "$words[1]" in
    start)
      compadd $(mux list) ;;
    stop)
      compadd $(mux list) ;;
  esac
}

_mux
