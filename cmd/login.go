package cmd

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/mattsre/flyhouse/pkg/config"
	"github.com/mattsre/flyhouse/pkg/log"
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

		fly.SetBaseURL(viper.GetString(config.ConfigFlyApiBase))
		auth, err := fly.StartCLISessionWebAuth(hostname, false)
		if err != nil {
			panic(err)
		}

		log.Println("Copy the following URL into a browser to continue: ", auth.URL)

		token, err := waitForCLISession(cmd.Context(), auth.ID)
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			log.Error("Login expired, please try again")
		case err != nil:
			log.Error("Failed to login, please try again", err)
		case token == "":
			log.Error("Failed to login, please try again")
		}

		viper.Set(config.ConfigFlyAccessToken, token)
		config.WriteViperConfig()

		user, err := fly.NewClientFromOptions(fly.ClientOptions{
			AccessToken: token,
		}).GetCurrentUser(context.TODO())

		if err != nil {
			log.Error("failed retrieving current user", err)
		}

		log.Println("Successfully logged in as ", user.Email)
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
			log.Errorf("failed retrieving token: %v\n", err)

			time.Sleep(time.Second)

			continue
		}

		break
	}

	return token, nil
}
