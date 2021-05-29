// Copyright (c) 2021, Satvik Reddy
// This is free software licensed under the BSD 2 Clause License

package main

import "fmt"

const (
	Black         = "\033[0;30m"
	Red           = "\033[0;31m"
	Green         = "\033[0;32m"
	Yellow        = "\033[0;33m"
	Blue          = "\033[0;34m"
	Magenta       = "\033[0;35m"
	Cyan          = "\033[0;36m"
	White         = "\033[0;37m"
	BrightBlack   = "\033[0;30;1m"
	BrightRed     = "\033[0;31;1m"
	BrightGreen   = "\033[0;32;1m"
	BrightYellow  = "\033[0;33;1m"
	BrightBlue    = "\033[0;34;1m"
	BrightMagenta = "\033[0;35;1m"
	BrightCyan    = "\033[0;36;1m"
	BrightWhite   = "\033[0;37;1m"
	Reset         = "\033[0;0m"
)

func startColor(color string) {
	fmt.Print(color)
}

func endColor() {
	fmt.Print(Reset)
}
