package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/spf13/cobra"

	"github.com/swampPr/minipt/internal/convert"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert an image from one format to the next",
	Example: `
	minipt convert --lossless --quality 75 ./in.png ./out.jpeg
	minipt convert ./in.jpeg ./out.webp
	`,
	RunE: execConvert,
}

var (
	quality  uint8
	lossless bool
)

func execConvert(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		_ = cmd.Help()
		return fmt.Errorf("please provide a path to the image and an output path")
	}

	if quality > 100 {
		quality = 100
	}

	absInPath, absOutPath, outFormat, err := checkPaths(args[0], args[1])
	if err != nil {
		_ = cmd.Help()
		return err
	}

	o := convert.Opts{
		Quality:  quality,
		Lossless: lossless,
	}

	fmt.Println(o)
	convertFn, ok := convert.Registry[outFormat]
	if !ok {
		return fmt.Errorf("invalid format")
	}

	err = convertFn(absInPath, absOutPath, o)
	if err != nil {
		return fmt.Errorf("there was an error converting the image: %w", err)
	}

	return nil
}

func checkPaths(inPath string, outPath string) (absInPath string, absOutPath string, outFormat string, err error) {
	outFormat = strings.TrimPrefix(filepath.Ext(outPath), ".")
	inFormat := strings.TrimPrefix(filepath.Ext(inPath), ".")
	supportedFormats := convert.SupportedFormats

	if !slices.Contains(supportedFormats, outFormat) || !slices.Contains(supportedFormats, inFormat) {
		return "", "", "", fmt.Errorf("invalid format: %s", outFormat)
	}

	_, err = os.Stat(filepath.Dir(outPath))
	if err != nil {
		return "", "", "", err
	}

	absOutPath, err = filepath.Abs(outPath)
	if err != nil {
		return "", "", "", err
	}

	absInPath, err = filepath.Abs(inPath)
	if err != nil {
		return "", "", "", err
	}

	_, err = os.Stat(absOutPath)
	if err == nil {
		return "", "", "", fmt.Errorf("please ensure the output path is not a directory or an existing file. ")
	}
	_, err = os.Stat(absInPath)
	if err != nil {
		return "", "", "", fmt.Errorf("please ensure the source image exists: %w", err)
	}

	return absInPath, absOutPath, outFormat, nil
}

func init() {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().Uint8VarP(&quality, "quality", "q", 90, "the quality level to use when encoding for webp and jpeg")
	convertCmd.Flags().BoolVarP(&lossless, "lossless", "l", false, "whether to encode lossless or not. default: NOT lossless")
}
