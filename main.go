package main

import (
	"fmt"

	"github.com/mattsre/flyhouse/cmd"
	"github.com/mattsre/flyhouse/pkg/config"
	"github.com/mattsre/flyhouse/pkg/log"
)

func main() {
	log.InitializeLogger()

	dir, err := config.GetConfigDirectory()
	if err != nil {
		log.Error("Error accessing home directory", err)
	}

	if err = config.InitConfigDir(dir); err != nil {
		log.Error(fmt.Sprintf("Error accessing config dir at %s", dir), err)
	}

	config.LoadViperConfig()

	root := cmd.RootCmd
	root.AddCommand(cmd.LoginCmd)
	root.AddCommand(cmd.DeployCmd)

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
