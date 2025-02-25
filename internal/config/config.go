package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var CfgFile string
var Token string

// initConfig reads in config file and ENV variables if set.
func InitConfig() {
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
		Token = viper.GetString("token")
	} else {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No config, but it's chill as long as you define ARTIFACTS_TOKEN in your env")
			Token = os.Getenv("ARTIFACTS_TOKEN")
			if Token == "" {
				fmt.Fprintf(os.Stderr, "Failed to locate config token file or environment variable, exiting")
				os.Exit(1)
			}
		} else {
			fmt.Println("failed loading config")
			os.Exit(1)
		}
	}
}
