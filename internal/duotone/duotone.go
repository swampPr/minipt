// Package duotone provides duotone  ->  Gives an image a nice duotone effect
package duotone

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

type Effect = string

const (
	Spotify       Effect = "spotify"
	Gryffindor    Effect = "gryffindor"
	Blueberry     Effect = "blueberry"
	MountainDew   Effect = "mountaindew"
	Yellow        Effect = "yellow"
	PinkSunset    Effect = "pinksunset"
	Clouds        Effect = "clouds"
	Seafoam       Effect = "seafoam"
	FireEngine    Effect = "fireengine"
	ElectricMelon Effect = "electricmelon"
	PurpleLight   Effect = "purplelight"
	Seashell      Effect = "seashell"
	Forrest       Effect = "forrest"
	Western       Effect = "western"
	IvyLeague     Effect = "ivyleague"
)

type Tones struct {
	Dark   color.RGBA
	Bright color.RGBA
}

var ErrEffect = errors.New("this is not a valid duotone effect")

var toneRegistry = map[Effect]Tones{
	Spotify:       {Dark: color.RGBA{R: 0, G: 0, B: 126}, Bright: color.RGBA{R: 106, G: 255, B: 127}},
	Gryffindor:    {Dark: color.RGBA{R: 104, G: 34, B: 24}, Bright: color.RGBA{R: 248, G: 190, B: 61}},
	Blueberry:     {Dark: color.RGBA{R: 127, G: 1, B: 211}, Bright: color.RGBA{R: 1, G: 219, B: 254}},
	MountainDew:   {Dark: color.RGBA{R: 1, G: 171, B: 109}, Bright: color.RGBA{R: 251, G: 240, B: 25}},
	Yellow:        {Dark: color.RGBA{R: 255, G: 93, B: 119}, Bright: color.RGBA{R: 251, G: 205, B: 32}},
	PinkSunset:    {Dark: color.RGBA{R: 17, G: 36, B: 94}, Bright: color.RGBA{R: 220, G: 67, B: 121}},
	Clouds:        {Dark: color.RGBA{R: 145, G: 207, B: 248}, Bright: color.RGBA{R: 255, G: 255, B: 255}},
	Seafoam:       {Dark: color.RGBA{R: 96, G: 36, B: 87}, Bright: color.RGBA{R: 172, G: 212, B: 157}},
	FireEngine:    {Dark: color.RGBA{R: 10, G: 5, B: 5}, Bright: color.RGBA{R: 240, G: 14, B: 46}},
	ElectricMelon: {Dark: color.RGBA{R: 80, G: 98, B: 214}, Bright: color.RGBA{R: 239, G: 0, B: 158}},
	PurpleLight:   {Dark: color.RGBA{R: 134, G: 130, B: 217}, Bright: color.RGBA{R: 222, G: 252, B: 254}},
	Seashell:      {Dark: color.RGBA{R: 101, G: 183, B: 214}, Bright: color.RGBA{R: 253, G: 217, B: 226}},
	Forrest:       {Dark: color.RGBA{R: 36, G: 26, B: 95}, Bright: color.RGBA{R: 1, G: 171, B: 109}},
	Western:       {Dark: color.RGBA{R: 54, G: 32, B: 12}, Bright: color.RGBA{R: 255, G: 151, B: 56}},
	IvyLeague:     {Dark: color.RGBA{R: 47, G: 7, B: 129}, Bright: color.RGBA{R: 223, G: 178, B: 51}},
}

func DuotoneImage(srcPath string, dstPath string, effect Effect) error {
	var logging logger.Logger
	tone, ok := toneRegistry[effect]
	if !ok {
		return ErrEffect
	}

	src, format, err := utils.LoadImage(srcPath)
	if err != nil {
		return err
	}

	inStat, err := os.Stat(srcPath)
	if err != nil {
		return err
	}

	srcName := strings.TrimSuffix(inStat.Name(), filepath.Ext(srcPath))
	outFormat := strings.TrimPrefix(filepath.Ext(dstPath), ".")

	dst := image.NewRGBA(src.Bounds())
	dstBounds := dst.Bounds()

	for y := dstBounds.Min.Y; y < dstBounds.Max.Y; y++ {
		for x := dstBounds.Min.X; x < dstBounds.Max.X; x++ {
			r, g, b, a := src.At(x, y).RGBA()

			r8, g8, b8, a8 := byte(r>>8), byte(g>>8), byte(b>>8), byte(a>>8) //#nosec

			luminosity := (0.2125 * float64(r8)) + (0.7152 * float64(g8)) + (0.0722 * float64(b8))
			t := luminosity / 255.0

			nR := float64(tone.Dark.R) + t*(float64(tone.Bright.R)-float64(tone.Dark.R))
			nG := float64(tone.Dark.G) + t*(float64(tone.Bright.G)-float64(tone.Dark.G))
			nB := float64(tone.Dark.B) + t*(float64(tone.Bright.B)-float64(tone.Dark.B))

			//#nosec
			newPix := color.RGBA{
				R: byte(nR),
				G: byte(nG),
				B: byte(nB),
				A: a8,
			}

			dst.Set(x, y, newPix)
		}
	}

	outStat, err := os.Stat(dstPath)
	if err == nil {
		if outStat.IsDir() {
			newName := fmt.Sprintf("%s_blur", srcName)
			err = utils.WriteImage(dst, format, newName, dstPath)
			if err != nil {
				return err
			}
			logging.Info("Image has been successfully written to: %s", dstPath)

			return nil
		}

		return fmt.Errorf("the output file may already exist. delete, rename or choose another out path")
	}

	if outFormat == "" {
		return fmt.Errorf("output path does not exist or is ambiguous e.g './output/file' ")
	}

	newName := strings.TrimSuffix(filepath.Base(dstPath), "."+outFormat)
	err = utils.WriteImage(dst, outFormat, newName, filepath.Dir(dstPath))
	if err != nil {
		return err
	}
	logging.Info("Successfully wrote image into: %s", dstPath)

	return nil
}
