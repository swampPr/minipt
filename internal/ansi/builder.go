// Package ansi provides ansi  INFO:  ANSI art business logic
package ansi

import (
	"fmt"

	"github.com/swampPr/minipt/internal/utils"
)

// Build function  INFO:  Builds the ANSI art
func Build(path string) error {
	src, _, err := utils.LoadImage(path)
	if err != nil {
		return err
	}

	dst, err := utils.ResizeToTerm(src)
	if err != nil {
		return err
	}

	for y := 0; y < dst.Bounds().Dy(); y += 2 {
		for x := range dst.Bounds().Dx() {
			topR, topG, topB, _ := dst.At(x, y).RGBA()
			bottomR, bottomG, bottomB, _ := dst.At(x, y+1).RGBA()

			topR8, topG8, topB8 := topR>>8, topG>>8, topB>>8
			bottomR8, bottomG8, bottomB8 := bottomR>>8, bottomG>>8, bottomB>>8

			fmt.Printf(
				"\033[38;2;%d;%d;%dm\033[48;2;%d;%d;%dm▀",
				topR8, topG8, topB8,
				bottomR8, bottomG8, bottomB8,
			)

		}
		fmt.Print("\033[0m\n")

	}

	return nil
}
