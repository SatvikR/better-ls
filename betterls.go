// Copyright (c) 2021, Satvik Reddy
// This is free software licensed under the BSD 2 Clause License

package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

const (
	Black         = "\u001b[30m"
	Red           = "\u001b[31m"
	Green         = "\u001b[32m"
	Yellow        = "\u001b[33m"
	Blue          = "\u001b[34m"
	Magenta       = "\u001b[35m"
	Cyan          = "\u001b[36m"
	White         = "\u001b[37m"
	BrightBlack   = "\u001b[30;1m"
	BrightRed     = "\u001b[31;1m"
	BrightGreen   = "\u001b[32;1m"
	BrightYellow  = "\u001b[33;1m"
	BrightBlue    = "\u001b[34;1m"
	BrightMagenta = "\u001b[35;1m"
	BrightCyan    = "\u001b[36;1m"
	BrightWhite   = "\u001b[37;1m"
	Reset         = "\u001b[0m"
)

var sizeLen int

func startColor(color string) {
	fmt.Print(color)
}

func endColor() {
	fmt.Print(Reset)
}

// printFile will properly printout file info given a file object
func printFile(file *fs.FileInfo) {
	startColor(BrightCyan)
	fmt.Printf("%s ", GetPermissionString(file))
	endColor()

	startColor(Magenta)
	for i := 0; i < (sizeLen - len(fmt.Sprint((*file).Size()))); i++ {
		fmt.Print(" ")
	}
	endColor()

	startColor(BrightBlue)
	fmt.Printf("%d ", (*file).Size())
	endColor()

	startColor(Green)
	fmt.Printf("%s\n", (*file).Name())
	endColor()
}

func main() {
	dir, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}

	dir = dir[:]

	sizeLen = GetSizeStringLen(dir)

	for _, file := range dir {
		printFile(&file)
	}
}
