// Copyright (c) 2021, Satvik Reddy
// This is free software licensed under the BSD 2 Clause License

package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

const (
	readBits    uint32 = 0b100
	writeBits   uint32 = 0b010
	executeBits uint32 = 0b001
)

// printFile will properly printout file info given a file object
func printFile(file *fs.FileInfo) {
	permissions := []rune("----------")

	mode := uint32((*file).Mode())

	if (*file).IsDir() {
		permissions[0] = 'd'
	}

	// Fill in permission based on mode bits
	for i := 2; i > -1; i-- {
		if mode&(readBits<<(i*3)) != 0 {
			permissions[i*3+1] = 'r'
		}

		if mode&(writeBits<<(i*3)) != 0 {
			permissions[i*3+2] = 'w'
		}

		if mode&(executeBits<<(i*3)) != 0 {
			permissions[i*3+3] = 'x'
		}
	}

	fmt.Printf("%s %s\n", string(permissions), (*file).Name())
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
