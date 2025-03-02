package embed

import (
	_ "embed"
)

//go:embed config.schema.json
var ConfigSchema string

//go:embed mux.zsh
var MuxZsh string
