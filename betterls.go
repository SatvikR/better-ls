// Copyright (c) 2021, Satvik Reddy
// This is free software licensed under the BSD 2 Clause License

package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

// printFile will properly printout file info given a file object
func printFile(file *fs.FileInfo) {
	fmt.Printf("%s %s\n", GetPermissionString(file), (*file).Name())
}

func main() {
	dir, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, file := range dir {
		printFile(&file)
	}
}
