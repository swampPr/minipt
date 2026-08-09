package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/swampPr/minipt/internal/ascii"
)

// asciiCmd represents the ascii command
var asciiCmd = &cobra.Command{
	Use:   "ascii",
	Short: "Turn your image into some nice ASCII art.",
	Example: `
	minipt ascii [IMAGE SOURCE PATH]
	`,
	RunE: execASCII,
}

func execASCII(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		_ = cmd.Help()
		return fmt.Errorf("you must provide a path to the source image")
	}

	absPath, err := filepath.Abs(args[0])
	if err != nil {
		return fmt.Errorf("an error occurred: %v", err)
	}
	err = ascii.Build(absPath)
	if err != nil {
		return fmt.Errorf("an error occurred trying to build the ascii art: %w", err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(asciiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// asciiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// asciiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
