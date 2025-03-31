/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/IAmRasputin/artifacts-go/internal/client"
	"github.com/IAmRasputin/artifacts-go/pkg/characters"
	"github.com/spf13/cobra"
)

// charactersCmd represents the characters command
var charactersCmd = &cobra.Command{
	Use:   "characters",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		artifactsClient, err := client.NewClientWithResponses(client.BaseURL)

		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to create artifacts client: %v", err)
			os.Exit(1)
		}

		charactersClient, err := characters.NewCharacterClient(*artifactsClient)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to create internal SDK client: %v", err)
			os.Exit(1)
		}

		chars, err := charactersClient.GetCharacters()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to fetch characters: %v", err)
			os.Exit(1)
		}

		fmt.Printf("%#v\n", chars)
	},
}

func init() {
	rootCmd.AddCommand(charactersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// charactersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// charactersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
