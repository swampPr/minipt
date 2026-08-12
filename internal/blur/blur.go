// Package blur provides blur  INFO: Image blurring business logic
package blur

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

// Opts struct  INFO:  Options for blurring
type Opts struct {
	Gaussian bool
	Passes   uint
}

func gaussian(src image.Image, x, y int) color.RGBA {
	var sumR, sumG, sumB, sumA, weightedSum int

	kernel := [3][3]int{
		{1, 2, 1},
		{2, 4, 2},
		{1, 2, 1},
	}

	for dy := -1; dy < 2; dy++ {
		for dx := -1; dx < 2; dx++ {
			weight := kernel[dy+1][dx+1]

			r, g, b, a := src.At(x+dx, y+dy).RGBA()

			sumR += int(r>>8) * weight
			sumG += int(g>>8) * weight
			sumB += int(b>>8) * weight
			sumA += int(a>>8) * weight

			weightedSum += weight
		}
	}

	//#nosec
	return color.RGBA{
		R: byte(sumR / weightedSum),
		G: byte(sumG / weightedSum),
		B: byte(sumB / weightedSum),
		A: byte(sumA / weightedSum),
	}
}

func box(src image.Image, x, y int) color.RGBA {
	var sumR, sumG, sumB, sumA int

	for dy := -1; dy < 2; dy++ {
		for dx := -1; dx < 2; dx++ {
			r, g, b, a := src.At(x+dx, y+dy).RGBA()

			sumR += int(r >> 8)
			sumG += int(g >> 8)
			sumB += int(b >> 8)
			sumA += int(a >> 8)
		}
	}

	return color.RGBA{
		R: byte(sumR / 9),
		G: byte(sumG / 9),
		B: byte(sumB / 9),
		A: byte(sumA / 9),
	}
}

func blur(src image.Image, o Opts) *image.RGBA {
	bounds := src.Bounds()

	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y + 1; y < bounds.Max.Y-1; y++ {
		for x := bounds.Min.X + 1; x < bounds.Max.X-1; x++ {
			if !o.Gaussian {
				dst.Set(x, y, box(src, x, y))
				continue
			}
			dst.Set(x, y, gaussian(src, x, y))
		}
	}

	return dst
}

// Apply function  INFO:  Applies the blur to the image given the blur options struct
func Apply(path string, o Opts, outPath string) error {
	var logging logger.Logger

	src, format, err := utils.LoadImage(path)
	if err != nil {
		return err
	}

	stat, err := os.Stat(path)
	if err != nil {
		return err
	}

	srcName := strings.TrimSuffix(stat.Name(), fmt.Sprintf(".%s", format))

	dst := blur(src, o)

	if o.Passes > 1 {
		for range o.Passes - 1 {
			dst = blur(dst, o)
		}
	}

	_, err = os.Stat(filepath.Dir(outPath))
	if err != nil {
		return err
	}

	var newName string
	outFormat := strings.Trim(filepath.Ext(outPath), ".")
	outStat, err := os.Stat(outPath)
	if err == nil {
		if outStat.IsDir() {
			newName = fmt.Sprintf("%s_blur", srcName)
			err = utils.WriteImage(dst, format, newName, outPath)
			if err != nil {
				return err
			}
			logging.Info("Successfully blurred image and placed it in: %s", outPath)
			return nil
		}
		return fmt.Errorf("output file already exists. delete, rename or choose another output path")
	}

	if outFormat == "" {
		return fmt.Errorf("output path does not exist or is ambiguous e.g './output/file' ")
	}

	newName = strings.TrimSuffix(filepath.Base(outPath), "."+outFormat)
	err = utils.WriteImage(dst, outFormat, newName, filepath.Dir(outPath))
	if err != nil {
		return err
	}
	logging.Info("Successfully blurred image and placed it in: %s", outPath)

	return nil
}
