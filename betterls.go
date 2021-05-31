// better-ls - A better version of ls
//
// Copyright (c) 2021, Satvik Reddy
// This is free software licensed under the BSD 2 Clause License

package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
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
func printFile(file fs.FileInfo, dirPath string) error {
	startColor(brightCyan)
	fmt.Printf("%s ", getPermissionString(file))
	endColor()

	startColor(brightYellow)
	owner, err := getFileOwner(file, dirPath)
	if err != nil {
		return err
	}
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

	return nil
}

func printFiles() error {
	dirPath := flag.Arg(0)
	if dirPath == "" {
		dirPath = "."
	}

	stat, err := os.Stat(dirPath)
	if err != nil {
		return err
	}
	mode := stat.Mode()

	if !mode.IsDir() {
		return errors.New("not a directory")
	}
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	absPath, err := filepath.Abs(dirPath)
	if err != nil {
		return err
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
		err := printFile(file, absPath)
		if err != nil {
			return err
		}
	}

	newLine()

	return nil
}

func main() {
	log.SetPrefix("better-ls: ")
	log.SetFlags(0)

	flag.Parse()

	err := printFiles()
	if err != nil {
		startColor(red)
		log.Print(err)
		endColor()
	}
}
