package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
	Use: "emailverifier",
	Short: "A CLI tool to verify email addresses",
	Long: "EmailVerifier is a simple CLI application to validate email addresses.",
	Run: func(cmd *cobra.Command, args []string)  {
		fmt.Println("Welcome to Email Verifier CLI. Use --help for commands.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error: ",err)
	}
}