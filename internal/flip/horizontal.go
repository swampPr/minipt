// Package flip provides flip  INFO:  Flipping an image vertically or horizontally
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

// Horizontal function  INFO:  Flips a given image horizontally
func Horizontal(inPath string, outPath string) error {
	var logging logger.Logger
	src, format, err := utils.LoadImage(inPath)
	if err != nil {
		return err
	}

	bounds := src.Bounds()
	outFormat := strings.TrimPrefix(filepath.Ext(outPath), ".")
	_, err = os.Stat(filepath.Dir(outPath))
	if err != nil {
		return err
	}

	var newName string

	srcName := strings.TrimSuffix(filepath.Base(inPath), "."+format)

	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			dst.Set(x, y, src.At(bounds.Dx()-1-x, y))
		}
	}

	outStat, err := os.Stat(outPath)
	if err == nil {
		if outStat.IsDir() {
			newName = fmt.Sprintf("%s_h_flip", srcName)

			err = utils.WriteImage(dst, format, newName, outPath)
			if err != nil {
				return err
			}

			logging.Info("Successfully wrote image to: %s", outPath)
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

	return nil
}
