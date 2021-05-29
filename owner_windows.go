// Copyright (c) 2021, Satvik Reddy
// This is free software licensed under the BSD 2 Clause License

package main

import (
	"io/fs"
	"path"

	"golang.org/x/sys/windows"
)

func getFileOwner(file fs.FileInfo, dirPath string) string {

	sd, err := windows.GetNamedSecurityInfo(
		path.Join(dirPath, file.Name()),
		windows.SE_FILE_OBJECT,
		windows.OWNER_SECURITY_INFORMATION,
	)
	if err != nil {
		panic(err)
	}

	ownerSID, _, err := sd.Owner()
	if err != nil {
		panic(err)
	}

	// Uses "" as system to look at the local system
	account, _, _, err := ownerSID.LookupAccount("")
	if err != nil {
		panic(err)
	}

	return account
}
