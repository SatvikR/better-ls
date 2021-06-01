// better-ls - A better version of ls
// Copyright (c) 2021, Satvik Reddy
//
// This is free software licensed under the BSD 2 Clause License

package main

import (
	"fmt"
	"sort"
	"strconv"
)

const (
	sizes string = "kMGTPE"
)

type flagsT struct {
	groupDirsFirst *bool
	humanReadable  *bool
}

type byDirectory []file

func (s byDirectory) Len() int {
	return len(s)
}

func (s byDirectory) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byDirectory) Less(i, j int) bool {
	return s[i].fsHandle.IsDir()
}

func handleFlags(files []file, flags flagsT) error {
	if *flags.groupDirsFirst {
		sort.Sort(byDirectory(files))
	}

	if *flags.humanReadable {
		for i, file := range files {
			hReadable, err := convertToHumanReadable(file.size)
			if err != nil {
				return err
			}

			files[i].size = hReadable
		}
	}

	return nil
}

func convertToHumanReadable(size string) (string, error) {
	const unit = 1000
	_b, err := strconv.Atoi(size)
	if err != nil {
		return "", err
	}

	b := int64(_b)

	if b < unit {
		return fmt.Sprintf("%d  B", b), nil
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), sizes[exp]), nil
}
