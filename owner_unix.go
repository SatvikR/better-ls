// Copyright (c) 2021, Satvik Reddy
// This is free software licensed under the BSD 2 Clause License

// +build aix darwin dragonfly freebsd js,wasm linux netbsd openbsd solaris

package main

func getFileOwner(file fs.FileInfo) (string, string) {
	panic("unimplemented")
}
