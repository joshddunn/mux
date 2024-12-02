package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

func Get() Config {
	//exhaustruct:ignore
	config := Config{}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	dir := fmt.Sprintf("%s/.mux.json", homeDir)
	file, err := os.ReadFile(dir)
	if err != nil {
		log.Fatal("Config file not found")
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Invalid config file")
	}

	err = config.Validate()
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func (config *Config) Validate() error {
	var err error

	for sessionIndex := range config.Sessions {
		session := &config.Sessions[sessionIndex]
		e := session.validate()
		err = errors.Join(err, e)

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
