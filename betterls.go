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
	fmt.Printf("%s ", GetPermissionString(file))

	for i := 0; i < (sizeLen - len(fmt.Sprint((*file).Size()))); i++ {
		fmt.Print(" ")
	}

	fmt.Printf("%d ", (*file).Size())
	fmt.Printf("%s\n", (*file).Name())
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
