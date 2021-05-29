// Copyright (c) 2021, Satvik Reddy
// This is free software licensed under the BSD 2 Clause License

// +build aix darwin dragonfly freebsd js,wasm linux netbsd openbsd solaris

package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"path"
	"syscall"
)

func getFileOwner(file fs.FileInfo, dirPath string) string {
	info, err := os.Stat(path.Join(dirPath, file.Name()))
	if err != nil {
		panic(err)
	}

	var uid int

	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		uid = int(stat.Uid)
	} else {
		panic("unable to get uid")
	}

	user, err := user.LookupId(fmt.Sprint(uid))
	if err != nil {
		panic(err)
	}

	return user.Name

}
