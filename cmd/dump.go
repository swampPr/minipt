package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/swampPr/minipt/internal/dump"
)

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump the entire image's bitmap to stdout or to a text file",
	Example: `
	minipt dump [IMAGE SRC]
	minipt dump -o  [IMAGE SRC] [DESTINATION FILE PATH]
	`,
	RunE: execDump,
}

func execDump(cmd *cobra.Command, args []string) error {
	var opts dump.Opts

	out, err := cmd.Flags().GetBool("out")
	if err != nil {
		return fmt.Errorf("an error occurred: %v", err)
	}

	if out {
		if len(args) < 2 || strings.TrimSpace(args[1]) == "" || strings.TrimSpace(args[0]) == "" {
			_ = cmd.Help()
			return fmt.Errorf("you must provide a path to the text file you want the bitmap saved to, and the path to the image")
		}

		outFileAbs, err := filepath.Abs(args[1])
		if err != nil {
			return fmt.Errorf("an error occurred: %v", err)
		}

		srcFileAbs, err := filepath.Abs(args[0])
		if err != nil {
			return fmt.Errorf("an error occurred: %v", err)
		}

		opts.OutFile = outFileAbs
		dump.BMP(opts, srcFileAbs)

		return nil
	}

	opts.OutFile = ""

	if len(args) < 1 || strings.TrimSpace(args[0]) == "" {
		return fmt.Errorf("you must provide a path to the source image")
	}
	srcFileAbs, err := filepath.Abs(args[0])
	if err != nil {
		return fmt.Errorf("an error occurred: %v", err)
	}

	dump.BMP(opts, srcFileAbs)

	return nil
}

func init() {
	rootCmd.AddCommand(dumpCmd)

	dumpCmd.Flags().BoolP("out", "o", false, "Whether to dump the bitmap into a text file")
}
