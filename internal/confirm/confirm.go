// Package confirm provides confirm  INFO:  Confirms image information
package confirm

import (
	"fmt"
	"image"
	"os"

	"github.com/swampPr/minipt/internal/logger"
)

// IMGInfo function  INFO:  Confirms and outputs image information
func IMGInfo(path string) error {
	var logging logger.Logger

	stat, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("an error occurred: %v", err)
	}

	f, err := os.Open(path) //#nosec
	if err != nil {
		return fmt.Errorf("an error occurred trying to open the file: %v", err)
	}

	_, format, err := image.Decode(f)
	if err != nil {
		return fmt.Errorf("an error occurred trying to decode the image: %v", err)
	}

	logging.Info("Actual Format: %s\n", format)
	logging.Info("Size: %d bytes\n", stat.Size())
	logging.Info("Permissions: %v\n", stat.Mode().Perm())
	logging.Info("ModTime: %v\n", stat.ModTime())

	return nil
}
