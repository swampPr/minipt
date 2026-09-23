package filter

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"os"
	"path/filepath"
	"strings"

	"github.com/swampPr/minipt/internal/logger"
	"github.com/swampPr/minipt/internal/utils"
)

func SepiaImage(inPath string, o *Opts) error {
	var logging logger.Logger

	src, format, err := utils.LoadImage(inPath)
	if err != nil {
		return err
	}

	inStat, err := os.Stat(inPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			logging.Error(false, "file does not exist: %s", inPath)
		}
		return err
	}

	srcName := strings.TrimSuffix(inStat.Name(), "."+format)
	outFormat := strings.TrimPrefix(filepath.Ext(o.OutPath), ".")

	//TODO: continue
	dst := image.NewRGBA(src.Bounds())
	for y := dst.Bounds().Min.Y; y < dst.Bounds().Max.Y; y++ {
		for x := dst.Bounds().Min.X; x < dst.Bounds().Max.X; x++ {
			r, g, b, a := src.At(x, y).RGBA()
			rf, gf, bf, a8 := float64(r>>8), float64(g>>8), float64(b>>8), byte(a>>8)

			nR := (rf * 0.393) + (gf * 0.769) + (bf * 0.189)
			nG := (rf * 0.349) + (gf * 0.686) + (bf * 0.168)
			nB := (rf * 0.272) + (gf * 0.534) + (bf * 0.131)

			cR, cG, cB := byte(max(0, min(nR, 255))), byte(max(0, min(nG, 255))), byte(max(0, min(nB, 255)))

			newPix := color.RGBA{
				R: cR,
				G: cG,
				B: cB,
				A: a8,
			}

			dst.Set(x, y, newPix)

		}
	}

	outStat, err := os.Stat(o.OutPath)
	if err == nil {
		if outStat.IsDir() {
			newName := fmt.Sprintf("%s_sepia", srcName)
			err := utils.WriteImage(dst, outFormat, newName, o.OutPath)
			if err != nil {
				return err
			}

			logging.Info("successfully wrote new image to: %s", filepath.Join(o.OutPath, newName))
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

	logging.Info("successfully wrote new image to: %s", o.OutPath)
	return nil
}
