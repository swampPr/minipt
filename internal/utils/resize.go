package utils

import (
	"image"
	"os"

	"golang.org/x/term"
)

// ResizeToTerm function  INFO:  Resizes a given image to the terminal output
func ResizeToTerm(src image.Image) (*image.RGBA, error) {
	bounds := src.Bounds()

	srcHeight := bounds.Dy()
	srcWidth := bounds.Dx()

	termWidth, termHeight, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return nil, err
	}

	targetWidth := termWidth - 2
	targetHeight := (termHeight - 2) * 2

	scaleX := float64(targetWidth) / float64(srcWidth)
	scaleY := float64(targetHeight) / float64(srcHeight)

	scale := min(scaleX, scaleY)

	newWidth := int(float64(srcWidth) * scale)
	newHeight := int(float64(srcHeight) * scale)

	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	for y := 0; y < dst.Bounds().Dy(); y++ {
		for x := 0; x < dst.Bounds().Dx(); x++ {
			srcX := x * srcWidth / newWidth
			srcY := y * srcHeight / newHeight

			dst.Set(x, y, src.At(srcX, srcY))
		}
	}

	return dst, nil
}
