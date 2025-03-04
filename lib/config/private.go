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
		session.SelectWindow = helpers.Pointer(1)
		if *session.ZeroIndex {
			session.SelectWindow = helpers.Pointer(0)
		}
	}

	minWindowIndex := 1
	maxWindowIndex := len(session.Windows)
	if *session.ZeroIndex {
		minWindowIndex--
		maxWindowIndex--
	}

	if *session.SelectWindow < minWindowIndex || *session.SelectWindow > maxWindowIndex {
		message := fmt.Sprintf("selectWindow must be between %d and %d.", minWindowIndex, maxWindowIndex)
		err = errors.Join(err, errors.New(message))
	}

	if !helpers.DirectoryExists(session.Dir) {
		err = errors.Join(err, directoryNotFoundError(session.Dir))
	}

	return err
}

func (window *Window) validate(session Session) error {
	var err error

	if window.Dir == "" {
		window.Dir = session.Dir
	} else if !helpers.DirectoryExists(window.Dir) {
		err = errors.Join(err, directoryNotFoundError(window.Dir))
	}

	if window.Layout == nil {
		window.Layout = helpers.Pointer(Default)
	}

	if window.SplitPercent == nil {
		window.SplitPercent = helpers.Pointer(35)
	}

	return err
}

func (pane *Pane) validate(window Window) error {
	var err error

	if pane.Dir == "" {
		pane.Dir = window.Dir
	} else if !helpers.DirectoryExists(pane.Dir) {
		err = errors.Join(err, directoryNotFoundError(pane.Dir))
	}

	if pane.Execute == nil {
		pane.Execute = helpers.Pointer(true)
	}

	return err
}

func directoryNotFoundError(dir string) error {
	message := fmt.Sprintf("Directory `%s` does not exist.", dir)
	return errors.New(message)
}
