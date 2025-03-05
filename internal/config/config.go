package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var CfgFile string
var token string

type TokenGetter interface {
	GetToken() (string, error)
}

// initConfig reads in config file and ENV variables if set.
func InitConfig() error {
	if CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		conf := home + "/.config/artifacts/"

		// Search config in home directory with name ".artifactsmmo-go" (without extension).
		viper.AddConfigPath(conf)
		viper.SetConfigName("token")
	}

	viper.SetEnvPrefix("artifacts")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		token = viper.GetString("token")
	} else {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No config, but it's chill as long as you define ARTIFACTS_TOKEN in your env")
			token = os.Getenv("ARTIFACTS_TOKEN")
			if token == "" {
				return fmt.Errorf("failed to locate config token file or environment variable")
			}
		} else {
			return fmt.Errorf("failed to load config")
		}
	}

	return nil
}

type DefaultTokenGetter struct {
	token string
}

func (d *DefaultTokenGetter) GetToken() (string, error) {
	if d.token == "" {
		err := InitConfig()

		if err != nil {
			return "", err
		}

		d.token = token
	}

	return d.token, nil
}

func NewDefaultTokenGetter() *DefaultTokenGetter {
	return &DefaultTokenGetter{
		token: token,
	}
}
