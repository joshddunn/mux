package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"mux/embed"
	"mux/lib/helpers"
	"os"
	"os/exec"
	"syscall"

	"github.com/xeipuuv/gojsonschema"
)

func ConfigDir() string {
	return fmt.Sprintf("%s/%s", helpers.HomeDir(), File)
}

func Get() Config {
	//exhaustruct:ignore
	config := Config{}

	file, err := os.ReadFile(ConfigDir())
	if err != nil {
		message := fmt.Sprintf("Config file (~/%s) does not exist. Run `mux config` and setup a config file.", File)
		log.Fatal(message)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		message := fmt.Sprintf("Config file (~/%s) cannot be parsed.", File)
		log.Fatal(message)
	}

	err = config.Validate()
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func (config *Config) Validate() error {
	var err error

	schemaLoader := gojsonschema.NewStringLoader(embed.ConfigSchema)
	configLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%s", ConfigDir()))

	result, err := gojsonschema.Validate(schemaLoader, configLoader)
	if err != nil {
		panic(err)
	}

	if !result.Valid() {
		for _, e := range result.Errors() {
			message := fmt.Sprintf("%s %s", e.Field(), e.Description())
			err = errors.Join(err, errors.New(message))
		}
	}

	sessionNames := map[string]bool{}

	for sessionIndex := range config.Sessions {
		session := &config.Sessions[sessionIndex]
		e := session.validate()
		err = errors.Join(err, e)

		if sessionNames[session.Name] {
			message := fmt.Sprintf("Session name must be unique. `%s` is repeated.", session.Name)
			err = errors.Join(err, errors.New(message))
		} else {
			sessionNames[session.Name] = true
		}

		for windowIndex := range session.Windows {
			window := &session.Windows[windowIndex]
			e := window.validate(*session)
			err = errors.Join(err, e)

			for paneIndex := range window.Panes {
				pane := &window.Panes[paneIndex]
				e := pane.validate(*window)
				err = errors.Join(err, e)
			}
		}
	}

	return err
}

func EditConfig() {
	editor := os.Getenv("EDITOR")

	binary, err := exec.LookPath(editor)
	if err != nil {
		panic(err)
	}

	args := append([]string{editor}, ConfigDir())
	err = syscall.Exec(binary, args, os.Environ())
	if err != nil {
		panic(err)
	}
}
