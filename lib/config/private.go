package config

import (
	"errors"
	"mux/lib/helpers"
)

func (session *Session) validate() error {
	var err error

	if session.Name == "" {
		err = errors.Join(err, errors.New("Session: Name is required"))
	}

	if session.Dir == "" {
		err = errors.Join(err, errors.New("Session: Dir is required"))
	} else if !helpers.DirectoryExists(session.Dir) {
		err = errors.Join(err, errors.New("Session: Invalid directory"))
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

func (window *Window) validate(session Session) error {
	var err error

	if window.Dir == "" {
		window.Dir = session.Dir
	}

	if !helpers.DirectoryExists(window.Dir) {
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

func (pane *Pane) validate(window Window) error {
	var err error

	if pane.Dir == "" {
		pane.Dir = window.Dir
	}

	if !helpers.DirectoryExists(pane.Dir) {
		err = errors.Join(err, errors.New("Pane: Invalid directory"))
	}

	if pane.Execute == nil {
		pane.Execute = helpers.Pointer(true)
	}

	return err
}
