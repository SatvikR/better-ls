// Copyright (c) 2021, Satvik Reddy
// This is free software licensed under the BSD 2 Clause License

package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

var sizeLen int

// printFile will properly printout file info given a file object
func printFile(file *fs.FileInfo) {
	startColor(BrightCyan)
	fmt.Printf("%s ", GetPermissionString(file))
	endColor()

	// Align the sizes properly
	for i := 0; i < (sizeLen - len(fmt.Sprint((*file).Size()))); i++ {
		fmt.Print(" ")
	}

	startColor(Magenta)
	fmt.Printf("%d ", (*file).Size())
	endColor()

	startColor(Red)
	fmt.Printf("%s ", (*file).ModTime().Format("Jan 02 15:04"))
	endColor()

	startColor(BrightGreen)
	fmt.Printf("%s\n", (*file).Name())
	endColor()
}

func main() {
	fmt.Println("")

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
