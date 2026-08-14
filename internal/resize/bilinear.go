// Package resize provides resize  INFO:  Handles resizing images using different algorithms
package resize

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"path/filepath"
	"strings"

	"github.com/swampPr/minipt/internal/utils"
)

func getFinalColor(src image.Image, dstBounds image.Rectangle, x, y int) color.RGBA {
	srcWidth := src.Bounds().Dx()
	srcHeight := src.Bounds().Dy()

	scaleX := float64(srcWidth) / float64(dstBounds.Dx())
	scaleY := float64(srcHeight) / float64(dstBounds.Dx())

	sourceX := float64(x) * scaleX
	sourceY := float64(y) * scaleY

	dx := sourceX - math.Floor(sourceX)
	dy := sourceY - math.Floor(sourceY)

	tlR, tlG, tlB, tlA := src.At(int(math.Floor(sourceX)), int(math.Floor(sourceY))).RGBA()
	trR, trG, trB, trA := src.At(int(math.Ceil(sourceX)), int(math.Floor(sourceY))).RGBA()

	blR, blG, blB, blA := src.At(int(math.Floor(sourceX)), int(math.Ceil(sourceY))).RGBA()
	brR, brG, brB, brA := src.At(int(math.Ceil(sourceX)), int(math.Ceil(sourceY))).RGBA()

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

	return nil
}
