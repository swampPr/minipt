package convert

import (
	"fmt"
	"os"

	"golang.org/x/image/bmp"

	"github.com/swampPr/minipt/internal/utils"
)

// BMP function  INFO:  Encodes the source image as a bitmap
func BMP(inPath string, outPath string, _ Opts) error {
	src, _, err := utils.LoadImage(inPath)
	if err != nil {
		return fmt.Errorf("there was an error loading the source image: %w", err)
	}

	newF, err := os.Create(outPath) //#nosec
	if err != nil {
		return fmt.Errorf("there was an error creating the new image: %w", err)
	}

	defer func() { _ = newF.Close() }()

	err = bmp.Encode(newF, src)
	if err != nil {
		return fmt.Errorf("there was an error encoding the new image: %w", err)
	}

	return nil
}
