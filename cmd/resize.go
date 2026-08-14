package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/swampPr/minipt/internal/resize"
)

var resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "Resize an image to a given width or height.",
	RunE:  execResize,
}

var (
	bilinear        bool
	nearestNeighbor bool
	width           uint
	height          uint
)

func execResize(_ *cobra.Command, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("you must provide a source path to the image and an out path")
	}

	if width <= 0 || height <= 0 {
		return fmt.Errorf("width and height must be a positive number bigger than 0")
	}

	absInPath, err := filepath.Abs(args[0])
	if err != nil {
		return err
	}

	absOutPath, err := filepath.Abs(args[1])
	if err != nil {
		return err
	}

	_, err = os.Stat(filepath.Dir(absOutPath))
	if err != nil {
		return err
	}

	_, err = os.Stat(absOutPath)
	if err == nil {
		return fmt.Errorf("ensure the outpath is not an existing file or directory. delete, rename or choose another outpath")
	}

	if bilinear {
		err = resize.Bilinear(absInPath, absOutPath, width, height)
		if err != nil {
			return err
		}

		return nil
	}

	err = resize.Nearest(absInPath, absOutPath, width, height)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(resizeCmd)

	resizeCmd.Flags().BoolVar(&bilinear, "bilinear", false, "Whether to use Bilinear interpolation in the resizing algorithm.")
	resizeCmd.Flags().BoolVar(&nearestNeighbor, "nearest", false, "Whether to use the nearest neighbor resizing algorithm.")
	resizeCmd.Flags().UintVarP(&width, "width", "x", 0, "The width to resize the image to.")
	resizeCmd.Flags().UintVarP(&height, "height", "y", 0, "The height to resize the image to.")

	resizeCmd.MarkFlagsMutuallyExclusive("bilinear", "nearest")
	resizeCmd.MarkFlagsOneRequired("bilinear", "nearest")

	err := resizeCmd.MarkFlagRequired("width")
	if err != nil {
		log.Fatal(err)
	}

	err = resizeCmd.MarkFlagRequired("height")
	if err != nil {
		log.Fatal(err)
	}
}
