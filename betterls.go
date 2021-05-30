// better-ls - A better version of ls
//
// Copyright (c) 2021, Satvik Reddy
// This is free software licensed under the BSD 2 Clause License

package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

var sizeLen int

func printFileName(file fs.FileInfo) {
	if file.IsDir() {
		startColor(brightYellow)
	} else {
		startColor(brightGreen)
	}

	fmt.Printf("%s", file.Name())

	endColor()
}

// printFile will properly printout file info given a file object
func printFile(file fs.FileInfo, dirPath string) {
	startColor(brightCyan)
	fmt.Printf("%s ", getPermissionString(file))
	endColor()

	startColor(brightYellow)
	owner := getFileOwner(file, dirPath)
	fmt.Printf("%s ", owner)
	endColor()

	// Align the sizes properly
	for i := 0; i < (sizeLen - len(fmt.Sprint(file.Size()))); i++ {
		fmt.Print(" ")
	}

	startColor(magenta)
	fmt.Printf("%d ", file.Size())
	endColor()

	startColor(red)
	fmt.Printf("%s ", file.ModTime().Format("Jan 02 15:04"))
	endColor()

	printFileName(file)

	newLine()
}

func printFiles(dirPath string) {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	absPath, err := filepath.Abs(dirPath)
	if err != nil {
		panic(err)
	}

	newLine()

	dir = dir[:]

	startColor(brightBlue)
	fmt.Printf("Directory: %s", absPath)
	endColor()
	newLine()

	sizeLen = getSizeStringLen(dir)

	newLine()

	for _, file := range dir {
		printFile(file, absPath)
	}

	newLine()

}

func main() {
	flag.Parse()

	dirPath := flag.Arg(0)
	if dirPath == "" {
		dirPath = "."
	}

	stat, err := os.Stat(dirPath)
	if err != nil {
		panic(err)
	}
	mode := stat.Mode()

	if !mode.IsDir() {
		panic("Not a directory")
	}

	printFiles(dirPath)
}
