package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"mux/lib/helpers"
	"os"
)

type Layout string

const (
	Default Layout = "default"
	Columns Layout = "columns"
	Rows    Layout = "rows"
)

type Config struct {
	Sessions []Session
}

type Session struct {
	Name         string   `json:"name"`
	Dir          string   `json:"dir"`
	ZeroIndex    *bool    `json:"zeroIndex"`
	SelectWindow *int     `json:"selectWindow"`
	Windows      []Window `json:"windows"`
}

type Window struct {
	Name         string  `json:"name"`
	Dir          string  `json:"dir"`
	Layout       *Layout `json:"layout"`
	SplitPercent *int    `json:"splitPercent"`
	Panes        []Pane  `json:"panes"`
}

type Pane struct {
	Dir     string `json:"dir"`
	Command string `json:"command"`
	Execute *bool  `json:"execute"`
}

func Get() (Config, error) {
	//exhaustruct:ignore
	config := Config{}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
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
	return config, err
}

func (config *Config) Validate() error {
	var err error

	for sessionIndex := range config.Sessions {
		session := &config.Sessions[sessionIndex]
		e := session.Validate()
		err = errors.Join(err, e)

		for windowIndex := range session.Windows {
			window := &session.Windows[windowIndex]
			e := window.Validate(*session)
			err = errors.Join(err, e)

			for paneIndex := range window.Panes {
				pane := &window.Panes[paneIndex]
				e := pane.Validate(*window)
				err = errors.Join(err, e)
			}
		}
	}

	return err
}

func (session *Session) Validate() error {
	var err error

	if session.Name == "" {
		err = errors.Join(err, errors.New("Session: Name is required"))
	}

	if session.Dir == "" {
		err = errors.Join(err, errors.New("Session: Dir is required"))
	} else {
		exists := helpers.DirectoryExists(session.Dir)
		if !exists {
			err = errors.Join(err, errors.New("Session: Invalid directory"))
		}
	}

	if session.ZeroIndex == nil {
		session.ZeroIndex = helpers.Pointer(false)
	}

	if session.SelectWindow == nil {
		if *session.ZeroIndex {
			session.SelectWindow = helpers.Pointer(0)
		} else {
			session.SelectWindow = helpers.Pointer(1)
		}
	} else if *session.ZeroIndex && (*session.SelectWindow < 0 || *session.SelectWindow > len(session.Windows)-1) {
		err = errors.Join(err, errors.New("Session: Invalid selectWindow value"))
	} else if !*session.ZeroIndex && (*session.SelectWindow < 1 || *session.SelectWindow > len(session.Windows)) {
		err = errors.Join(err, errors.New("Session: Invalid selectWindow value"))
	}

	if len(session.Windows) == 0 {
		err = errors.Join(err, errors.New("Session: A session must have at least one window"))
	}

	return err
}

func (window *Window) Validate(session Session) error {
	var err error

	if window.Dir == "" {
		window.Dir = session.Dir
	}

	exists := helpers.DirectoryExists(window.Dir)
	if !exists {
		err = errors.Join(err, errors.New("Window: Invalid directory"))
	}

	if window.Layout == nil {
		window.Layout = helpers.Pointer(Default)
	} else {
		switch *window.Layout {
		case Default, Columns, Rows:
		// do nothing
		default:
			err = errors.Join(err, errors.New("Window: Invalid layout"))
		}
	}

	if window.Name == "" {
		err = errors.Join(err, errors.New("Window: Name is required"))
	}

	if window.SplitPercent == nil {
		if *window.Layout == Default {
			window.SplitPercent = helpers.Pointer(35)
		}
	} else if *window.Layout != Default {
		err = errors.Join(err, errors.New("Window: Split percent can only be defined if using the default layout"))
	} else if *window.SplitPercent <= 0 || *window.SplitPercent >= 100 {
		err = errors.Join(err, errors.New("Window: Invalid split percent"))
	}

	return err
}

func (pane *Pane) Validate(window Window) error {
	var err error

	if pane.Dir == "" {
		pane.Dir = window.Dir
	}

	exists := helpers.DirectoryExists(pane.Dir)
	if !exists {
		err = errors.Join(err, errors.New("Pane: Invalid directory"))
	}

	if pane.Execute == nil {
		pane.Execute = helpers.Pointer(true)
	}

	return err
}
