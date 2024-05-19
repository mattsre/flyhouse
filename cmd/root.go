package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "flyhouse",
	Short: "Flyhouse is a CLI for deploying Clickhouse clusters on Fly.io",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
