// Copyright (c) 2021, Satvik Reddy
// This is free software licensed under the BSD 2 Clause License

package main

import (
	"io/fs"

	"golang.org/x/sys/windows"
)

func getFileOwner(file fs.FileInfo) (string, string) {

	sd, err := windows.GetNamedSecurityInfo(
		file.Name(),
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

	groupSID, _, err := sd.Owner()
	if err != nil {
		panic(err)
	}

	// Uses "" as system to look at the local system
	account, _, _, err := ownerSID.LookupAccount("")
	if err != nil {
		panic(err)
	}

	group, _, _, err := groupSID.LookupAccount("")
	if err != nil {
		panic(err)
	}

	return account, group
}
