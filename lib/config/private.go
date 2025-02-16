package config

import (
	"errors"
	"fmt"
	"mux/lib/helpers"
)

func (session *Session) validate() error {
	var err error

	if session.ZeroIndex == nil {
		session.ZeroIndex = helpers.Pointer(false)
	}

	if session.SelectWindow == nil {
		if *session.ZeroIndex {
			session.SelectWindow = helpers.Pointer(0)
		} else {
			session.SelectWindow = helpers.Pointer(1)
		}
	}

	if !helpers.DirectoryExists(session.Dir) {
		message := fmt.Sprintf("%s does not exist", session.Dir)
		err = errors.Join(err, errors.New(message))
	}

	return err
}

func (window *Window) validate(session Session) error {
	var err error

	if window.Dir == "" {
		window.Dir = session.Dir
	} else if !helpers.DirectoryExists(window.Dir) {
		message := fmt.Sprintf("%s does not exist", window.Dir)
		err = errors.Join(err, errors.New(message))
	}

	if window.Layout == nil {
		window.Layout = helpers.Pointer(Default)
	}

	if window.SplitPercent == nil && *window.Layout == Default {
		window.SplitPercent = helpers.Pointer(35)
	}

	return err
}

func (pane *Pane) validate(window Window) error {
	var err error

	if pane.Dir == "" {
		pane.Dir = window.Dir
	} else if !helpers.DirectoryExists(pane.Dir) {
		message := fmt.Sprintf("%s does not exist", pane.Dir)
		err = errors.Join(err, errors.New(message))
	}

	if pane.Execute == nil {
		pane.Execute = helpers.Pointer(true)
	}

	return err
}
