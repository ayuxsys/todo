package cmd

import (
	"os"
	sqlite3 "todo/internal/database/sqlite"
	"todo/pkg/config"

	"github.com/spf13/cobra"
)

func (c *cmdi) LoadCmd() *cobra.Command {
	loadCmd := &cobra.Command{
		Use:   "load",
		Short: "load a config to the database",
		Run:   c.runLoad,
	}
	loadCmd.PersistentFlags().StringVarP(&c.load.cfg, "config", "c", defaultCfgPath, "path to the config file to load")
	loadCmd.PersistentFlags().BoolVarP(&c.load.removeDB, "remove", "r", false, "remove previously created database before loading config (recommended)")
	return loadCmd
}

func (c *cmdi) runLoad(_ *cobra.Command, _ []string) {
	if c.load.cfg != "" {
		if c.load.removeDB {
			os.Remove(c.database)
		}
		conn := sqlite3.Connect(sqlite3.Dsn(c.database))
		defer conn.Close()
		cfg := config.Load(c.load.cfg)

		if err := conn.InitTables(); err != nil {
			panic(err)
		}

		if err := conn.InsertTaks(cfg.Tasks); err != nil {
			panic(err)
		}
	}
}
