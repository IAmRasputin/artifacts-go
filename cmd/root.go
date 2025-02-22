/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// $ af
var rootCmd = &cobra.Command{
	Use:   "af",
	Short: "A Go client for the Artifacts MMORPG",
	Long: `A Go client and command-line interface for playing and interact with
Artifacts, an API-based MMORPG.`,
	Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/artifacts/token.yaml")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
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
		// TODO
	} else {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No config, but it's chill as long as you define ARTIFACTS_TOKEN in your env")
		} else {
			fmt.Println("failed loading config")
			os.Exit(1)
		}
	}
}
