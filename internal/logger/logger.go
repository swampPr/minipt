// Package logger provides logger  INFO:  Provides logging information and functionality
package logger

import (
	"os"

	"github.com/fatih/color"
)

// Logger struct  INFO:  The logger struct used to create a new logger instance
type Logger struct{}

// NewColor function  INFO:  Returns a new color that can be used to color strings
func NewColor(r, g, b int) func(a ...any) string {
	return color.RGB(r, g, b).SprintFunc()
}

// Info method  INFO:  Prints with a green text
func (Logger) Info(format string, a ...any) {
	color.Green(format, a...)
}

// Error method  INFO:  Prints with a red text
func (Logger) Error(exit bool, format string, a ...any) {
	color.Red(format, a...)
	if exit {
		os.Exit(1)
	}
}

// Warn method  INFO:  Prints with an orange text
func (Logger) Warn(format string, a ...any) {
	color.RGB(255, 165, 0).Printf(format, a...)
}
