package resize

import (
	"fmt"
	"image"
	"path/filepath"
	"strings"

	"github.com/swampPr/minipt/internal/utils"
)

// Nearest function  INFO:  Resizes an iamge usisng the nearerst neighbor algorithm
func Nearest(inPath, outPath string, width, height uint) error {
	src, _, err := utils.LoadImage(inPath)
	if err != nil {
		return err
	}

	outFormat := strings.TrimPrefix(filepath.Ext(outPath), ".")
	if outFormat == "" {
		return fmt.Errorf("outpath is ambiguous: e.g './output/file'")
	}

	srcWidth := src.Bounds().Dx()
	srcHeight := src.Bounds().Dy()

	dst := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))

	for y := range dst.Bounds().Dx() {
		for x := range dst.Bounds().Dy() {
			sourceX := (x * srcWidth) / int(width)
			sourceY := (y * srcHeight) / int(height)

			dst.Set(x, y, src.At(sourceX, sourceY))
		}
	}

	outName := strings.TrimSuffix(filepath.Base(outPath), "."+outFormat)

	err = utils.WriteImage(dst, outFormat, outName, filepath.Dir(outPath))
	if err != nil {
		return err
	}

	return nil
}
