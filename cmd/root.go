// Package cmd provides cmd  INFO:  Root command "minipt"
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/swampPr/minipt/internal/logger"
)

var rootCmd = &cobra.Command{
	Use:   "minipt",
	Short: "A simple image processing tool.",
	Run:   executeCmd,
}

func executeCmd(_ *cobra.Command, _ []string) {
	log := logger.Logger{}
	log.Error(true, "Please enter a subcommand/action. Use 'minipt --help' for more information")
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	rootCmd.SilenceUsage = true

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
