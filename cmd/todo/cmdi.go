package cmd

import "os"

type cmdi struct {
	database string // path to the sqlite database
	current  bool   // get current task
	load     load
}

type load struct {
	cfg      string
	removeDB bool
}

var (
	defaultDatabasePath = os.Getenv("HOME") + "/.todo/tasks.db"
	defaultCfgPath      = os.Getenv("HOME") + "/.todo/config.yaml"
)
