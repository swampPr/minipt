package ascii

import (
	"image"
	// "image/jpeg"  INFO:  setting jpeg and png decoder
	_ "image/jpeg"
	_ "image/png"
	"os"

	// "golang.org/x/image/bmp"  INFO:  setting bmp and webp decoder
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/webp"
)

func loadImage(path string) (image.Image, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path) //#nosec
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}
