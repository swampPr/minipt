package filter

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"path/filepath"
	"strings"

	"github.com/swampPr/minipt/internal/logger"
	"github.com/swampPr/minipt/internal/utils"
)

func applyInvert(src image.Image, x, y int) color.RGBA {
	r, g, b, a := src.At(x, y).RGBA()

	r8, g8, b8, a8 := byte(r>>8), byte(g>>8), byte(b>>8), byte(a>>8) //#nosec

	return color.RGBA{
		R: 255 - r8,
		G: 255 - g8,
		B: 255 - b8,
		A: a8,
	}
}

// Invert function  INFO:  Inverts the colors in an image
func Invert(inPath string, o *Opts) error {
	var logging logger.Logger

	src, format, err := utils.LoadImage(inPath)
	if err != nil {
		return err
	}

	inStat, err := os.Stat(inPath)
	if err != nil {
		return err
	}

	srcName := strings.TrimSuffix(inStat.Name(), "."+format)
	outFormat := strings.TrimPrefix(filepath.Ext(o.OutPath), ".")

	dst := image.NewRGBA(src.Bounds())

	for y := src.Bounds().Min.Y; y < src.Bounds().Max.Y; y++ {
		for x := src.Bounds().Min.X; x < src.Bounds().Max.X; x++ {
			dst.Set(x, y, applyInvert(src, x, y))
		}
	}

	_, err = os.Stat(filepath.Dir(o.OutPath))
	if err != nil {
		return err
	}

	outStat, err := os.Stat(o.OutPath)
	if err == nil {
		if outStat.IsDir() {
			newName := fmt.Sprintf("%s_blur", srcName)
			err = utils.WriteImage(dst, format, newName, o.OutPath)
			if err != nil {
				return err
			}

			logging.Info("Successfully wrote image into: %s", o.OutPath)

			return nil
		}

		return fmt.Errorf("the path already exists. delete, rename or choose another output path")
	}

	if outFormat == "" {
		return fmt.Errorf("output path is ambiguous or invalid: e.g './output/file'")
	}

	newName := strings.TrimSuffix(filepath.Base(o.OutPath), "."+outFormat)
	err = utils.WriteImage(dst, outFormat, newName, filepath.Dir(o.OutPath))
	if err != nil {
		return err
	}

	logging.Info("Successfully wrote image into: %s", o.OutPath)
	return nil
}
