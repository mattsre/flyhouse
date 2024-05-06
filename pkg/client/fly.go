package client

import (
	"context"

	"github.com/spf13/viper"
	"github.com/superfly/fly-go"

	"github.com/mattsre/flyhouse/pkg/config"
	"github.com/mattsre/flyhouse/pkg/log"
)

type FlyClient struct {
	client *fly.Client
}

func NewFlyClient() *FlyClient {
	baseUrl := viper.GetString(config.ConfigFlyApiBase)
	token := viper.GetString(config.ConfigFlyAccessToken)

	client := fly.NewClientFromOptions(fly.ClientOptions{
		BaseURL:     baseUrl,
		AccessToken: token,
	})

	return &FlyClient{
		client: client,
	}
}

func (fc *FlyClient) CreateApp(ctx context.Context, name string) {
	app, err := fc.client.CreateApp(ctx, fly.CreateAppInput{
		Name:            name,
		OrganizationID:  "matt-981",
		PreferredRegion: fly.StringPointer("bos"),
	})

	if err != nil {
		log.Error("failed creating app in Fly: ", err)
		return
	}

	log.Println("created app successfully: ", app.Name)
}
