package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/swampPr/minipt/internal/flip"
)

var flipCmd = &cobra.Command{
	Use:   "flip",
	Short: "Flip an image horizontally or vertically",
	RunE:  execFlip,
}

var (
	horizontal bool
	vertical   bool
)

func execFlip(_ *cobra.Command, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("please input a source path and output path")
	}

	absInPath, err := filepath.Abs(args[0])
	if err != nil {
		return err
	}

	absOutPath, err := filepath.Abs(args[1])
	if err != nil {
		return err
	}

	if horizontal {
		err = flip.Horizontal(absInPath, absOutPath)
		if err != nil {
			return err
		}
		return nil
	}

	err = flip.Vertical(absInPath, absOutPath)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(flipCmd)

	flipCmd.Flags().BoolVarP(&horizontal, "horizontal", "x", false, "Whether to flip an image horizontally")
	flipCmd.Flags().BoolVarP(&vertical, "vertical", "y", false, "Whether to flip an image vertically")

	flipCmd.MarkFlagsOneRequired("horizontal", "vertical")
	flipCmd.MarkFlagsMutuallyExclusive("horizontal", "vertical")
}
