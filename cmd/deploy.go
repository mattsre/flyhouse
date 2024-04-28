package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a Clickhouse cluster using the given configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create the cluster")
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
