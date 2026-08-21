package cmd

import (
	"fmt"
	sqlite3 "todo/internal/database/sqlite"
	"todo/internal/utils/timestamp"

	"github.com/spf13/cobra"
)

var Root *cobra.Command

func init() {
	cmdi := &cmdi{}
	Root = cmdi.Root()
	Root.SilenceErrors = true
}

func (c *cmdi) Root() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "todo",
		Run: c.runRoot,
	}
	rootCmd.PersistentFlags().StringVarP(&c.database, "database", "d", defaultDatabasePath, "path to the sqlite database")
	rootCmd.PersistentFlags().BoolVarP(&c.current, "task-json", "t", false, "print current task in json format to load in modules such as waybar")

	rootCmd.AddCommand(c.LoadCmd())

	return rootCmd
}

func (c *cmdi) runRoot(_ *cobra.Command, _ []string) {
	if c.current {
		conn := sqlite3.Connect(sqlite3.Dsn(c.database))
		defer conn.Close()
		task, err := conn.CurrentTask()
		if err != nil {
			panic(err)
		}

		text := task.Title + ": " + timestamp.ToHourMinute(task.Time.Start) + "-" + timestamp.ToHourMinute(task.Time.End)
		toolTip := task.Desc

		fmt.Printf(`{"text":"%s","tooltip":"%s"}`, text, toolTip)
	}
}
