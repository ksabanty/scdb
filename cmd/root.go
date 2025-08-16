package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scdb",
	Short: "Sports Cards DB manages your card collection via the CLI",
	Long:  `A CLI tool to manage a sports card collection with add, delete, and list commands.`,
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
