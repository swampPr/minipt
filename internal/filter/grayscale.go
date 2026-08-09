package filter

import "github.com/swampPr/minipt/internal/utils"

func Apply(path string, outPath string) error {
	src, err := utils.LoadImage(path)
	if err != nil {
		return err
	}

	return nil
}
