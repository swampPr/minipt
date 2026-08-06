// Package dump provides dump  INFO:  Business logic for dumping an image's bitmap
package dump

import (
	"bufio"
	"fmt"
	"image"
	"io"
	"os"

	_ "image/jpeg" // "image/jpeg"  INFO: Register jpeg encoder
	_ "image/png"

	_ "golang.org/x/image/bmp" // "golang.org/x/image/bmp"  INFO:  Register BMP encoder
	_ "golang.org/x/image/webp"

	"github.com/swampPr/minipt/internal/logger"
)

// Opts struct  INFO:  Optiosn for dumping bmp
type Opts struct {
	OutFile string
}

// BMP function  INFO:  Dumps the image's bitmap to stdout or to a text file
func BMP(o Opts, src string) error {
	var logging logger.Logger

	_, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("an error occurred: %v", err)
	}

	f, err := os.Open(src) //#nosec
	if err != nil {
		return fmt.Errorf("an error occurred opening the file: %v", err)
	}

	defer func() { _ = f.Close() }()

	img, _, err := image.Decode(f)
	if err != nil {
		return fmt.Errorf("an error occurred decoding the image: %v", err)
	}

	if o.OutFile == "" {
		logging.Info("RGB Bitmap:\n")
		writer := bufio.NewWriter(os.Stdout)
		pixels := imageToBytes(img)
		defer func() { _ = writer.Flush() }()
		writeBytes(pixels, writer)
		return nil
	}

	newFile, err := os.Create(o.OutFile)
	if err != nil {
		return fmt.Errorf("an error occurred creating the txt file: %v", err)
	}

	defer func() { _ = newFile.Close() }()

	writer := bufio.NewWriter(newFile)
	pixels := imageToBytes(img)
	writeBytes(pixels, writer)

	return nil
}

func imageToBytes(img image.Image) []byte {
	bounds := img.Bounds()

	pixels := make([]byte, 0, bounds.Dx()*bounds.Dy()*4)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()

			pixels = append(pixels, byte(r>>8), byte(g>>8), byte(b>>8), byte(a>>8)) //#nosec
		}
	}

	return pixels
}

func writeBytes(pixels []byte, w io.Writer) {
	for i, b := range pixels {
		if i > 0 {
			fmt.Fprintf(w, " ")
		}

		fmt.Fprintf(w, "%d", b)
	}
}
