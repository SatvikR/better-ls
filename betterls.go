// better-ls - A better version of ls
// Copyright (c) 2021, Satvik Reddy
//
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
	"path"
	"path/filepath"
)

var sizeLen int
var absDirPath string

type file struct {
	mode     string
	owner    string
	size     string
	time     string
	name     string
	fsHandle fs.FileInfo
}

func constructFiles(dir []fs.FileInfo) ([]file, error) {
	files := make([]file, len(dir))

	for i, fileInfo := range dir {
		owner, err := getFileOwner(fileInfo, absDirPath)
		if err != nil {
			return nil, err
		}

		var fileSize int64

		if fileInfo.IsDir() {
			fileSize = getDirSize(path.Join(absDirPath, fileInfo.Name()))
		} else {
			fileSize = fileInfo.Size()
		}

		files[i] = file{
			mode:     getPermissionString(fileInfo),
			owner:    owner,
			size:     fmt.Sprint(fileSize),
			time:     fileInfo.ModTime().Format("Jan 02 15:04"),
			name:     fileInfo.Name(),
			fsHandle: fileInfo,
		}
	}

	return files, nil
}

func printFileName(file file) {
	if file.fsHandle.IsDir() {
		startColor(brightYellow)
	} else {
		startColor(brightGreen)
	}

	fmt.Printf("%s", file.name)

	endColor()
}

// printFile will properly printout file info given a file object
func printFile(file file) error {
	startColor(brightCyan)
	fmt.Printf("%s ", file.mode)
	endColor()

	startColor(brightYellow)
	fmt.Printf("%s ", file.owner)
	endColor()

	// Align the sizes properly
	for i := 0; i < (sizeLen - len(file.size)); i++ {
		fmt.Print(" ")
	}

	startColor(magenta)
	fmt.Printf("%s ", file.size)
	endColor()

	startColor(red)
	fmt.Printf("%s ", file.time)
	endColor()

	printFileName(file)

	newLine()

	return nil
}

func printFiles(flags flagsT) error {
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

	absDirPath, err = filepath.Abs(dirPath)
	if err != nil {
		return err
	}

	dir = dir[:]

	files, err := constructFiles(dir)
	if err != nil {
		return err
	}

	err = handleFlags(files, flags)
	if err != nil {
		return err
	}

	newLine()

	startColor(brightBlue)
	fmt.Printf("Directory: %s", absDirPath)
	endColor()
	newLine()

	sizeLen = getSizeStringLen(files)

	newLine()

	for _, file := range files {
		err := printFile(file)
		if err != nil {
			return err
		}
	}

	newLine()

	return nil
}

func main() {
	flags := flagsT{}

	log.SetPrefix("better-ls: ")
	log.SetFlags(0)

	flags.groupDirsFirst = flag.Bool(
		"group-dirs-first",
		false,
		"output all the directories at the top",
	)

	flags.humanReadable = flag.Bool(
		"human-readable",
		false,
		"list sizes in a human readable format",
	)

	flag.Parse()

	err := printFiles(flags)
	if err != nil {
		startColor(red)
		log.Print(err)
		endColor()
	}
}
