// Copyright (c) 2021, Satvik Reddy
// This is free software licensed under the BSD 2 Clause License

package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

var sizeLen int

func printFileName(file *fs.FileInfo) {
	if (*file).IsDir() {
		startColor(brightYellow)
	} else {
		startColor(brightGreen)
	}

	fmt.Printf("%s", (*file).Name())

	endColor()
}

// printFile will properly printout file info given a file object
func printFile(file *fs.FileInfo) {
	startColor(brightCyan)
	fmt.Printf("%s ", getPermissionString(file))
	endColor()

	// Align the sizes properly
	for i := 0; i < (sizeLen - len(fmt.Sprint((*file).Size()))); i++ {
		fmt.Print(" ")
	}

	startColor(magenta)
	fmt.Printf("%d ", (*file).Size())
	endColor()

	startColor(red)
	fmt.Printf("%s ", (*file).ModTime().Format("Jan 02 15:04"))
	endColor()

	printFileName(file)

	newLine()
}

func main() {
	newLine()

	dir, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}

	dir = dir[:]

	sizeLen = getSizeStringLen(dir)

	for _, file := range dir {
		printFile(&file)
	}

	newLine()
}
