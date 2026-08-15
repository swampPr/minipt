package convert

import (
	"fmt"
	"os"

	"github.com/deepteams/webp"

	"github.com/swampPr/minipt/internal/logger"
	"github.com/swampPr/minipt/internal/utils"
)

// WebP function  INFO:  Converts a given image to WebP
func WebP(inPath string, outPath string, o Opts) error {
	var logging logger.Logger

	src, _, err := utils.LoadImage(inPath)
	if err != nil {
		return fmt.Errorf("there was an error loading the image: %w", err)
	}

	newF, err := os.Create(outPath) //#nosec
	if err != nil {
		return fmt.Errorf("there was an error creating the new image: %w", err)
	}

	err = webp.Encode(newF, src, &webp.EncoderOptions{Quality: float32(o.Quality), Lossless: o.Lossless})
	if err != nil {
		return fmt.Errorf("there was an error encoding the new image: %w", err)
	}

	logging.Info("Successfully wrote WEBP image into: %s", outPath)

	return nil
}
