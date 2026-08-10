package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/swampPr/minipt/internal/filter"
)

var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "Set a filter on your image",
	RunE:  execFilter,
}

var filterName string

var gray bool

func execFilter(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		_ = cmd.Help()
		fmt.Println(len(args))
		fmt.Println(args)
		return fmt.Errorf("please provide a path to the image and an output path")
	}

	inAbsPath, err := filepath.Abs(args[0])
	if err != nil {
		return fmt.Errorf("an error occurred: %s", err)
	}

	outAbsPath, err := filepath.Abs(args[1])
	if err != nil {
		return fmt.Errorf("an error occurred: %s", err)
	}

	switch {
	case gray:
		filterFn, ok := filter.Registry["grayscale"]
		if !ok {
			_ = cmd.Help()
			return fmt.Errorf("unknown filter: %s", filterName)
		}

		opts := filter.Opts{
			OutPath: outAbsPath,
		}

		err = filterFn(inAbsPath, &opts)
		if err != nil {
			return fmt.Errorf("an error occurred trying to set the filter: %w", err)
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(filterCmd)

	filterCmd.Flags().BoolVarP(&gray, "gray", "g", false, "Grayscale filter toggle.")

	filterCmd.MarkFlagsOneRequired("gray")
	filterCmd.MarkFlagsMutuallyExclusive("gray")
}
