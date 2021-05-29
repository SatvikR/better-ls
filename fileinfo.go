// Copyright (c) 2021, Satvik Reddy
// This is free software licensed under the BSD 2 Clause License

package main

import (
	"fmt"
	"io/fs"
)

const (
	readBits    uint32 = 0b100
	writeBits   uint32 = 0b010
	executeBits uint32 = 0b001
)

// Parses the permission bits into a string
func GetPermissionString(file *fs.FileInfo) string {
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

	return string(permissions)
}

func GetSizeStringLen(dir []fs.FileInfo) int {
	l := 0

	for _, f := range dir {
		sizeStringLen := len(fmt.Sprint(f.Size()))
		if sizeStringLen > l {
			l = sizeStringLen
		}
	}

	return l
}