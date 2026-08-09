package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/swampPr/minipt/internal/ansi"
)

var ansiCmd = &cobra.Command{
	Use:   "ansi",
	Short: "Turn your image into cute ANSI art",
	Example: `
	minipt ansi [SOURCE IMAGE]
	`,
	RunE: execAnsi,
}

func execAnsi(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		_ = cmd.Help()
		return fmt.Errorf("you must provide a path to the source image")
	}

	absPath, err := filepath.Abs(args[0])
	if err != nil {
		return fmt.Errorf("an error occurred: %v", err)
	}

	err = ansi.Build(absPath)
	if err != nil {
		return fmt.Errorf("an error occurred trying to build the ANSI art: %w", err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(ansiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ansiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ansiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
