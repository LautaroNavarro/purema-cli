package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "purema",
	Short: "Purema is a pull request manager",
	Long: "Purema helps you to manage the team's pull request. Tracking and reminding you when needed",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}