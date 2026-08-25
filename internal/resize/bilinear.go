// Package resize provides resize  INFO:  Handles resizing images using different algorithms
package resize

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"path/filepath"
	"strings"

	"github.com/swampPr/minipt/internal/logger"
	"github.com/swampPr/minipt/internal/utils"
)

func getFinalColor(src image.Image, dstBounds image.Rectangle, x, y int) color.RGBA {
	srcWidth := src.Bounds().Dx()
	srcHeight := src.Bounds().Dy()

	scaleX := float64(srcWidth-1) / float64(dstBounds.Dx()-1)
	scaleY := float64(srcHeight-1) / float64(dstBounds.Dy()-1)

	sourceX := float64(x) * scaleX
	sourceY := float64(y) * scaleY

	x0 := int(math.Floor(sourceX))
	x1 := min(x0+1, srcWidth-1)

	y0 := int(math.Floor(sourceY))
	y1 := min(y0+1, srcHeight-1)

	dx := sourceX - float64(x0)
	dy := sourceY - float64(y0)

	tlR, tlG, tlB, tlA := src.At(x0, y0).RGBA()
	trR, trG, trB, trA := src.At(x1, y0).RGBA()

	blR, blG, blB, blA := src.At(x0, y1).RGBA()
	brR, brG, brB, brA := src.At(x1, y1).RGBA()

	topR := (float64(tlR) * (1 - dx)) + (float64(trR) * dx)
	topG := (float64(tlG) * (1 - dx)) + (float64(trG) * dx)
	topB := (float64(tlB) * (1 - dx)) + (float64(trB) * dx)
	topA := (float64(tlA) * (1 - dx)) + (float64(trA) * dx)

	botR := (float64(blR) * (1 - dx)) + (float64(brR) * dx)
	botG := (float64(blG) * (1 - dx)) + (float64(brG) * dx)
	botB := (float64(blB) * (1 - dx)) + (float64(brB) * dx)
	botA := (float64(blA) * (1 - dx)) + (float64(brA) * dx)

	finalR := (topR * (1 - dy)) + (botR * dy)
	finalG := (topG * (1 - dy)) + (botG * dy)
	finalB := (topB * (1 - dy)) + (botB * dy)
	finalA := (topA * (1 - dy)) + (botA * dy)

	//#nosec
	return color.RGBA{
		R: byte(int(finalR) >> 8),
		G: byte(int(finalG) >> 8),
		B: byte(int(finalB) >> 8),
		A: byte(int(finalA) >> 8),
	}
}

// Bilinear function  INFO:  Uses bilinear interpolation to resize an image
func Bilinear(inPath string, outPath string, width uint, height uint) error {
	var logging logger.Logger

	src, _, err := utils.LoadImage(inPath)
	if err != nil {
		return err
	}

	outFormat := strings.TrimPrefix(filepath.Ext(outPath), ".")

	dst := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	for y := range dst.Bounds().Dy() {
		for x := range dst.Bounds().Dx() {
			dst.Set(x, y, getFinalColor(src, dst.Bounds(), x, y))
		}
	}

	if outFormat == "" {
		return fmt.Errorf("outpath is ambiguous: e.g './output/file'")
	}

	outName := strings.TrimSuffix(filepath.Base(outPath), "."+outFormat)

	err = utils.WriteImage(dst, outFormat, outName, filepath.Dir(outPath))
	if err != nil {
		return err
	}

	logging.Info("Successfully wrote image to: %s", outPath)

	return nil
}
