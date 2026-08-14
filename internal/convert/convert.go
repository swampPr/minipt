// Package convert provides convert  INFO:  Deals with converting images to other formats
package convert

// Opts struct  INFO:  Opts for compression. Only used for JPEG and WEBP funcs
type Opts struct {
	Quality  uint8
	Lossless bool
}

// SupportedFormats  INFO:  Lists the supported formats
var SupportedFormats = []string{"bmp", "webp", "png", "jpeg", "jpg"}

// Func  INFO:  Signature for a function that converts from one format to another
type Func = func(inPath string, outPath string, o Opts) error

// Registry  INFO: Registry of allowed formats to convert to and their functions
var Registry = map[string]Func{
	"jpeg": Jpeg,
	"jpg":  Jpeg,
	"webp": WebP,
	"png":  PNG,
	"bmp":  BMP,
}
