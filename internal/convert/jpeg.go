package convert

import (
	"fmt"
	"image/jpeg"
	"os"

	"github.com/swampPr/minipt/internal/logger"
	"github.com/swampPr/minipt/internal/utils"
)

// Jpeg function  INFO:  Converts the given image to jpeg
func Jpeg(inPath string, outPath string, o Opts) error {
	var logging logger.Logger

	img, _, err := utils.LoadImage(inPath)
	if err != nil {
		return fmt.Errorf("there was an issue loading the image: %w", err)
	}

	newF, err := os.Create(outPath) //#nosec
	if err != nil {
		return fmt.Errorf("there was a problem creating the new image: %w", err)
	}

	defer func() { _ = newF.Close() }()

	err = jpeg.Encode(newF, img, &jpeg.Options{Quality: int(o.Quality)})
	if err != nil {
		return fmt.Errorf("there was a problem encoding the new image: %w", err)
	}

	logging.Info("Successfully wrote JPEG file into: %s", outPath)

	return nil
}
