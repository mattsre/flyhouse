package cmd

import (
	"fmt"
	"os"

	"github.com/mattsre/flyhouse/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	fly "github.com/superfly/fly-go"
)

var rootCmd = &cobra.Command{
	Use:   "flyhouse",
	Short: "Flyhouse is a CLI for deploying Clickhouse clusters on Fly.io",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	dir, err := helpers.GetConfigDirectory()
	if err != nil {
		fmt.Println("Error accessing home directory", err)
	}

	if err = helpers.InitConfigDir(dir); err != nil {
		fmt.Println(fmt.Sprintf("Error accessing config dir at %s", dir), err)
	}

	helpers.LoadViperConfig()

	fly.SetBaseURL(viper.GetString(helpers.ConfigFlyApiBase))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
