// Copyright (c) 2021, Satvik Reddy
// This is free software licensed under the BSD 2 Clause License

package main

import "fmt"

const (
	black         = "\033[0;30m"
	red           = "\033[0;31m"
	green         = "\033[0;32m"
	yellow        = "\033[0;33m"
	blue          = "\033[0;34m"
	magenta       = "\033[0;35m"
	cyan          = "\033[0;36m"
	white         = "\033[0;37m"
	brightBlack   = "\033[0;30;1m"
	brightRed     = "\033[0;31;1m"
	brightGreen   = "\033[0;32;1m"
	brightYellow  = "\033[0;33;1m"
	brightBlue    = "\033[0;34;1m"
	brightMagenta = "\033[0;35;1m"
	brightCyan    = "\033[0;36;1m"
	brightWhite   = "\033[0;37;1m"
	reset         = "\033[0;0m"
)

func startColor(color string) {
	fmt.Print(color)
}

func endColor() {
	fmt.Print(reset)
}

func newLine() {
	fmt.Println("")
}
