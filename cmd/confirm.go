package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/swampPr/minipt/internal/confirm"
)

var confirmCmd = &cobra.Command{
	Use:   "confirm",
	Short: "Confirm and view an image's information such as format, size, permissions and other information.",
	Example: `
	minipt confirm [SOURCE IMAGE]
	`,
	RunE: execConfirm,
}

func execConfirm(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		_ = cmd.Help()
		return fmt.Errorf("you must provide a path to the source image")
	}

	absPath, err := filepath.Abs(args[0])
	if err != nil {
		return fmt.Errorf("an error occurred: %v", err)
	}

	err = confirm.IMGInfo(absPath)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(confirmCmd)
}
