// Package ascii provides ascii  INFO:  Builds the ascii art of the image
package ascii

import (
	"fmt"

	"github.com/swampPr/minipt/internal/utils"
)

const asciiChars = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/|()1{}[]?-_+~<>i!lI;:,\"^'\\`. "

type pixMap []byte

func grayToASCII(gray byte) string {
	index := (float64(gray) / float64(255)) * float64((len(asciiChars) - 1))
	return string(asciiChars[int(index)])
}

func (pm pixMap) reduce(fn func(ac string, cur byte, idx int) string, initVal string) string {
	ac := initVal

	for i, v := range pm {
		ac = fn(ac, v, i)
	}

	return ac
}

// Build function  INFO:  Builds the actual ascii string
func Build(path string) error {
	src, _, err := utils.LoadImage(path)
	if err != nil {
		return err
	}

	img, err := utils.ResizeImage(src)
	if err != nil {
		return err
	}

	rgb := make([]byte, 0, img.Bounds().Dx()*img.Bounds().Dy())

	for y := range img.Bounds().Dy() {
		for x := range img.Bounds().Dx() {
			r, g, b, _ := img.At(x, y).RGBA()

			rgb = append(rgb, byte(r>>8), byte(g>>8), byte(b>>8)) //#nosec
		}
	}

	grayScale := pixMap(colorToGray(rgb))

	ascii := grayScale.reduce(func(ac string, cur byte, idx int) string {
		nextChars := grayToASCII(cur)

		if (idx+1)%img.Bounds().Dx() == 0 {
			nextChars += "\n"
		}
		return ac + nextChars
	}, "")

	fmt.Printf("%v", ascii)

	return nil
}
