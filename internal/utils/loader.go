// Package utils provides utils  INFO:  Useful utilities
package utils

import (
	"image"
	// image/jpeg  INFO:  Set jpeg and png decoder
	_ "image/jpeg"
	_ "image/png"
	"os"

	// "github.com/deepteams/webp"  INFO:  set webp and bmp decoder
	_ "github.com/deepteams/webp"
	_ "golang.org/x/image/bmp"
)

// LoadImage function  INFO:  Loads the image and returns an image.Image object
func LoadImage(path string) (image.Image, string, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, "", err
	}

	f, err := os.Open(path) //#nosec
	if err != nil {
		return nil, "", err
	}

	defer func() { _ = f.Close() }()

	img, format, err := image.Decode(f)
	if err != nil {
		return nil, "", err
	}

	return img, format, nil
}
