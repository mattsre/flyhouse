package cmd

import (
	"fmt"

	"github.com/mattsre/flyhouse/pkg/config"
	"github.com/mattsre/flyhouse/pkg/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "flyhouse",
	Short: "Flyhouse is a CLI for deploying Clickhouse clusters on Fly.io",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	dir, err := config.GetConfigDirectory()
	if err != nil {
		log.Error("Error accessing home directory", err)
	}

	if err = config.InitConfigDir(dir); err != nil {
		log.Error(fmt.Sprintf("Error accessing config dir at %s", dir), err)
	}

	config.LoadViperConfig()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
