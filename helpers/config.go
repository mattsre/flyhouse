package helpers

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

const (
	ConfigFlyApiBase     = "fly_api_base"
	ConfigFlyAccessToken = "fly_access_token"
)

var writableConfigKeys = []string{ConfigFlyAccessToken}

func GetConfigDirectory() (string, error) {
	if value, isSet := os.LookupEnv("FLYHOUSE_CONFIG_DIR"); isSet {
		return value, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}

	return filepath.Join(home, ".config/flyhouse"), nil
}

func GetConfigPath() (string, error) {
	configDir, err := GetConfigDirectory()
	if err != nil {
		return "", err
	}

	return path.Join(configDir, "config.yml"), nil
}

func InitConfigDir(dir string) error {
	if !DirectoryExists(dir) {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return err
		}
	}

	return nil
}

func LoadViperConfig() error {
	configFilePath, err := GetConfigPath()
	if err != nil {
		return err
	}

	viper.SetConfigFile(configFilePath)
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err == nil {
		// TODO: change to debug log
		fmt.Println("loaded flyhouse config file from", viper.ConfigFileUsed())
	}

	if errors.Is(err, fs.ErrNotExist) {
		if err = WriteViperConfig(); err != nil {
			return err
		}
	}

	viper.SetDefault(ConfigFlyApiBase, "https://api.fly.io")

	return nil
}

func WriteViperConfig() error {
	configFilePath, err := GetConfigPath()
	if err != nil {
		return err
	}

	out := map[string]interface{}{}

	for k, v := range viper.AllSettings() {
		if persistConfigKey(k) {
			out[k] = v
		}
	}

	data, err := yaml.Marshal(&out)
	if err != nil {
		return err
	}

	return os.WriteFile(configFilePath, data, 0o600)
}

func persistConfigKey(key string) bool {
	if viper.InConfig(key) {
		return true
	}

	for _, k := range writableConfigKeys {
		if k == key {
			return true
		}
	}

	return false
}
