package convert

import (
	"fmt"
	"image/png"
	"os"

	"github.com/swampPr/minipt/internal/utils"
)

// PNG function  INFO:  encodes the source image into a png image
func PNG(inPath string, outPath string, _ Opts) error {
	src, _, err := utils.LoadImage(inPath)
	if err != nil {
		return fmt.Errorf("there was an error loading the source image: %w", err)
	}

	newF, err := os.Create(outPath) //#nosec
	if err != nil {
		return fmt.Errorf("there was an error creating the new image: %w", err)
	}

	defer func() { _ = newF.Close() }()

	err = png.Encode(newF, src)
	if err != nil {
		return fmt.Errorf("there was an error encoding the new image: %w", err)
	}

	return nil
}
