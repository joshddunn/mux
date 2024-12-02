package config

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

type Layout string
