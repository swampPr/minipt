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

func applyGrayscale(src image.Image, x, y int) color.RGBA {
	r, g, b, a := src.At(x, y).RGBA()

	r8, g8, b8, a8 := byte(r>>8), byte(g>>8), byte(b>>8), byte(a>>8) //#nosec

	grayScale := 0.2126*float64(r8) + 0.7152*float64(g8) + 0.0722*float64(b8)

	return color.RGBA{
		R: byte(grayScale),
		G: byte(grayScale),
		B: byte(grayScale),
		A: a8,
	}
}

// Grayscale function  INFO:  applies the gray scale to the image
func Grayscale(inPath string, o *Opts) error {
	var logging logger.Logger

	src, format, err := utils.LoadImage(inPath)
	if err != nil {
		return err
	}

	inStat, err := os.Stat(inPath)
	if err != nil {
		return err
	}

	_, err = os.Stat(filepath.Dir(o.OutPath))
	if err != nil {
		return fmt.Errorf("path does not exist")
	}

	srcName := strings.TrimSuffix(inStat.Name(), filepath.Ext(inPath))
	outFormat := strings.TrimPrefix(filepath.Ext(o.OutPath), ".")

	dst := image.NewRGBA(src.Bounds())

	for y := src.Bounds().Min.Y; y < src.Bounds().Max.Y; y++ {
		for x := src.Bounds().Min.X; x < src.Bounds().Max.X; x++ {
			dst.Set(x, y, applyGrayscale(src, x, y))
		}
	}

	outStat, err := os.Stat(o.OutPath)
	if err == nil {
		if outStat.IsDir() {
			newName := fmt.Sprintf("%s_blur", srcName)
			err = utils.WriteImage(dst, format, newName, o.OutPath)
			if err != nil {
				return err
			}
			logging.Info("Image has been successfully written to: %s", o.OutPath)

			return nil
		}

		return fmt.Errorf("the output file may already exist. delete, rename or choose another out path")
	}

	if outFormat == "" {
		return fmt.Errorf("output path does not exist or is ambiguous e.g './output/file' ")
	}

	newName := strings.TrimSuffix(filepath.Base(o.OutPath), "."+outFormat)
	err = utils.WriteImage(dst, outFormat, newName, filepath.Dir(o.OutPath))
	if err != nil {
		return err
	}
	logging.Info("Successfully wrote image into: %s", o.OutPath)
	return nil
}
