package cmd

import (
	"github.com/mattsre/flyhouse/pkg/client"
	"github.com/mattsre/flyhouse/pkg/log"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a Clickhouse cluster using the given configuration",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting deployment of flyhouse app!")

		client := client.NewFlyClient()

		client.CreateApp(cmd.Context(), "flyhouse-app")
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
