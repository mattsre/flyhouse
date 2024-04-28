package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/mattsre/flyhouse/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	fly "github.com/superfly/fly-go"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to your Fly.io account to grant API access",
	Run: func(cmd *cobra.Command, args []string) {
		hostname, err := os.Hostname()
		if err != nil {
			panic(err)
		}

		auth, err := fly.StartCLISessionWebAuth(hostname, false)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Copy the following URL into a browser to continue: %s\n", auth.URL)

		token, err := waitForCLISession(cmd.Context(), auth.ID)
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			fmt.Println("Login expired, please try again")
		case err != nil:
			fmt.Println("Failed to login, please try again", err)
		case token == "":
			fmt.Println("Failed to login, please try again")
		}

		viper.Set(helpers.ConfigFlyAccessToken, token)
		helpers.WriteViperConfig()

		fmt.Println("Logged in successfully!")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func waitForCLISession(parent context.Context, authId string) (token string, err error) {
	ctx, cancel := context.WithTimeout(parent, 15*time.Minute)
	defer cancel()

	for ctx.Err() == nil {
		if token, err = fly.GetAccessTokenForCLISession(ctx, authId); err != nil {
			fmt.Printf("failed retrieving token: %v\n", err)

			time.Sleep(time.Second)

			continue
		}

		break
	}

	return token, nil
}
