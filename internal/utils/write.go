package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"github.com/deepteams/webp"
	"golang.org/x/image/bmp"
)

// WriteImage function  INFO:  Writes the image out to the path provides
func WriteImage(dst image.Image, format string, name string, path string) error {
	fileName := fmt.Sprintf("%s.%s", name, format)

	joinedPaths := filepath.Join(path, fileName)

	absPath, err := filepath.Abs(joinedPaths)
	if err != nil {
		return err
	}

	f, err := os.Create(absPath) //#nosec
	if err != nil {
		return err
	}

	switch format {
	case "png":
		err := png.Encode(f, dst)
		if err != nil {
			return fmt.Errorf("an error occurred while encoding the new image: %w", err)
		}
	case "jpeg":
		err := jpeg.Encode(f, dst, nil)
		if err != nil {
			return fmt.Errorf("an error occurred while encoding the new image: %w", err)
		}

	case "jpg":
		err := jpeg.Encode(f, dst, nil)
		if err != nil {
			return fmt.Errorf("an error occurred while encoding the new image: %w", err)
		}

	case "bmp":
		err := bmp.Encode(f, dst)
		if err != nil {
			return fmt.Errorf("an error occurred while encoding the new image: %w", err)
		}
	case "webp":
		err := webp.Encode(f, dst, nil)
		if err != nil {
			return fmt.Errorf("an error occurred while encoding the new image: %w", err)
		}

	default:
		return fmt.Errorf("invalid format '%s'. please choose between png, jpeg and webp", format)
	}

	return nil
}
