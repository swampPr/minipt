package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/swampPr/minipt/internal/duotone"
)

var effect string

// duotoneCmd represents the duotone command
var duotoneCmd = &cobra.Command{
	Use:   "duotone",
	Short: "Give your image a nice duotone filter effect.",
	Example: `
	minipt duotone -l [to list all the duotone effects]
	minipt duotone -f blueberry ./in.png ./out.png
	`,
	RunE: execDuotone,
}

func execDuotone(cmd *cobra.Command, args []string) error {
	listFlag := cmd.Flags().Changed("list")

	if listFlag {
		fmt.Println(" All the available duotone effects: spotify, gryffindor, blueberry, mountaindew, yellow, pinksunset, clouds, seafoam, fireengine, electricmelon,\n purplelight, seashell, forrest, western, ivyleague ")
		return nil
	}

	if len(args) < 2 {
		_ = cmd.Help()
		return fmt.Errorf("please provide a path to the source image and an output path")
	}

	absInPath, err := filepath.Abs(args[0])
	if err != nil {
		return err
	}

	absOutPath, err := filepath.Abs(args[1])
	if err != nil {
		return err
	}

	return duotone.DuotoneImage(absInPath, absOutPath, duotone.Effect(strings.TrimSpace(effect)))
}

func init() {
	rootCmd.AddCommand(duotoneCmd)

	duotoneCmd.Flags().BoolP("list", "l", false, "List all duotone effects")
	duotoneCmd.Flags().StringVarP(&effect, "effect", "e", "", "The effect to use")
}
