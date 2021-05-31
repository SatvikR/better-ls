// better-ls - A better version of ls
// Copyright (c) 2021, Satvik Reddy
//
// This is free software licensed under the BSD 2 Clause License

package main

import (
	"io/fs"
	"sort"
)

type flagsT struct {
	groupDirsFirst *bool
}

type byDirectory []fs.FileInfo

func (s byDirectory) Len() int {
	return len(s)
}

func (s byDirectory) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byDirectory) Less(i, j int) bool {
	return s[i].IsDir()
}

func handleFlags(dir []fs.FileInfo, flags flagsT) {
	if *flags.groupDirsFirst {
		sort.Sort(byDirectory(dir))
	}
}
