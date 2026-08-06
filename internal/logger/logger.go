// Package logger provides logger  INFO:  Provides logging information and functionality
package logger

import (
	"os"

	"github.com/fatih/color"
)

type Logger struct{}

func NewColor(r, g, b int) func(a ...any) string {
	return color.RGB(r, g, b).SprintFunc()
}

func (Logger) Info(format string, a ...any) {
	color.Green(format, a...)
}

func (Logger) Error(exit bool, format string, a ...any) {
	color.Red(format, a...)
	if exit {
		os.Exit(1)
	}
}

func (Logger) Warn(format string, a ...any) {
	color.RGB(255, 165, 0).Printf(format, a...)
}
