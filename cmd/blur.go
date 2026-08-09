package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/swampPr/minipt/internal/blur"
)

var blurCmd = &cobra.Command{
	Use:   "blur",
	Short: "Blur your image and choose between gaussian or box blur",
	Example: `
	minipt blur [SOURCE IMAGE PATH]
	minipt blur -p=4 [SOURCE IMAGE PATH]
	minipt blur -g [SOURCE IMAGE PATH]
	minipt blur -g -p=4 [SOURCE IMAGE PATH]
	`,
	RunE: execBlur,
}

func execBlur(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		_ = cmd.Help()
		return fmt.Errorf("you must provide a path to the source image")
	}

	p, err := cmd.Flags().GetUint("passes")
	if err != nil {
		return fmt.Errorf("an error occurred: %w", err)
	}

	if p < 1 {
		_ = cmd.Help()
		return fmt.Errorf("you must provide at least 1 pass of blurring")
	}

	absPath, err := filepath.Abs(args[0])
	if err != nil {
		return fmt.Errorf("an error occurred: %w", err)
	}

	absOutPath, err := filepath.Abs(args[1])
	if err != nil {
		return fmt.Errorf("an error occurred: %w", err)
	}

	g, err := cmd.Flags().GetBool("gaussian")
	if err != nil {
		return fmt.Errorf("an error occurred: %w", err)
	}

	opts := blur.Opts{
		Gaussian: g,
		Passes:   p,
	}

	err = blur.Apply(absPath, opts, absOutPath)
	if err != nil {
		return fmt.Errorf("an error occurred trying to blur the image: %w", err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(blurCmd)

	blurCmd.Flags().BoolP("gaussian", "g", false, "Whether to use Gaussian blur")
	blurCmd.Flags().UintP("passes", "p", 1, "How many passes. The more passes the blurrier the image. Default 1")
}
