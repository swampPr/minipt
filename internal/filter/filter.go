// Package filter provides filter  INFO:  Filter package deals with adding filters to images
package filter

// Opts struct  INFO:  The options for filters
type Opts struct {
	OutPath string
}

// Func  INFO:  The signature of a function that adds a filter to an image
type Func func(path string, fo *Opts) error

// Registry  INFO:  Registry for available filters and their function references
var Registry = map[string]Func{
	"grayscale": Grayscale,
	"invert":    Invert,
}
