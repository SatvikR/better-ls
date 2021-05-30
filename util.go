// better-ls - A better version of ls
//
// Copyright (c) 2021, Satvik Reddy
// This is free software licensed under the BSD 2 Clause License

package main

import "fmt"

const (
	black         string = "\033[0;30m"
	red           string = "\033[0;31m"
	green         string = "\033[0;32m"
	yellow        string = "\033[0;33m"
	blue          string = "\033[0;34m"
	magenta       string = "\033[0;35m"
	cyan          string = "\033[0;36m"
	white         string = "\033[0;37m"
	brightBlack   string = "\033[0;30;1m"
	brightRed     string = "\033[0;31;1m"
	brightGreen   string = "\033[0;32;1m"
	brightYellow  string = "\033[0;33;1m"
	brightBlue    string = "\033[0;34;1m"
	brightMagenta string = "\033[0;35;1m"
	brightCyan    string = "\033[0;36;1m"
	brightWhite   string = "\033[0;37;1m"
	reset         string = "\033[0;0m"
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
