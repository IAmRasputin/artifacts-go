/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/IAmRasputin/artifacts-go/internal/client"
	"github.com/IAmRasputin/artifacts-go/pkg/status"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Gets the status of the game server",
	Run: func(cmd *cobra.Command, args []string) {
		artifactsClient, err := client.NewClientWithResponses(client.BaseURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create base game client: %s", err)
			os.Exit(1)
		}

		resp, err := status.NewGameStatusClient(artifactsClient).GetGameServerStatus()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get game server status: %s", err)
			os.Exit(1)
		}

		fmt.Printf("%#v\n", resp)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
