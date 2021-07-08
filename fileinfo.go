// better-ls - A better version of ls
// Copyright (c) 2021, Satvik Reddy
//
// This is free software licensed under the BSD 2 Clause License

package main

import (
	"io/fs"
)

const (
	readBits    uint32 = 0b100
	writeBits   uint32 = 0b010
	executeBits uint32 = 0b001
)

// Parses the permission bits into a string
func getPermissionString(file fs.FileInfo) string {
	permissions := []rune("----------")

	mode := uint32(file.Mode())

	if file.IsDir() {
		permissions[0] = 'd'
	}

	// Fill in permission based on mode bits
	for i := 2; i > -1; i-- {
		if mode&(readBits<<(i*3)) != 0 {
			permissions[(3-i-1)*3+1] = 'r'
		}

		if mode&(writeBits<<(i*3)) != 0 {
			permissions[(3-i-1)*3+2] = 'w'
		}

		if mode&(executeBits<<(i*3)) != 0 {
			permissions[(3-i-1)*3+3] = 'x'
		}
	}

	return string(permissions)
}

func getSizeStringLen(files []file) int {
	l := 0

	for _, f := range files {
		sizeStringLen := len(f.size)
		if sizeStringLen > l {
			l = sizeStringLen
		}
	}

	return l
}

func getDirSize(path string) int64 {
	return 0
}
