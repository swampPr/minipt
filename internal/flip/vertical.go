package flip

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strings"

	"github.com/swampPr/minipt/internal/logger"
	"github.com/swampPr/minipt/internal/utils"
)

// Vertical function  INFO:  Flips image vertically
func Vertical(inPath string, outPath string) error {
	var logging logger.Logger

	src, format, err := utils.LoadImage(inPath)
	if err != nil {
		return err
	}

	bounds := src.Bounds()
	outFormat := strings.TrimPrefix(filepath.Ext(outPath), ".")

	srcName := strings.TrimSuffix(filepath.Base(inPath), "."+format)

	_, err = os.Stat(filepath.Dir(outPath))
	if err != nil {
		return err
	}

	var newName string

	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			dst.Set(x, y, src.At(x, bounds.Dy()-1-y))
		}
	}

	outStat, err := os.Stat(outPath)
	if err == nil {
		if outStat.IsDir() {
			newName = fmt.Sprintf("%s_v_flip", srcName)

			err = utils.WriteImage(dst, format, newName, outPath)
			if err != nil {
				return err
			}

			logging.Info("Successfully wrote image into: %s", outPath)
			return nil
		}

		return fmt.Errorf("ensure the output file does not exist. delete, rename or choose another outpath")
	}

	if outFormat == "" {
		return fmt.Errorf("output path does not exist or is ambiguous e.g './output/file' ")
	}

	newName = strings.TrimSuffix(filepath.Base(outPath), "."+outFormat)
	err = utils.WriteImage(dst, outFormat, newName, filepath.Dir(outPath))
	if err != nil {
		return fmt.Errorf("there was an error writing the new image: %w", err)
	}

	logging.Info("Successfully wrote image to: %s", outPath)

	return nil
}
