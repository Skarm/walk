// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package walk

import (
	"github.com/Skarm/win"
	"syscall"
)

func CreateMutex(lpMutexAttributes *win.SECURITY_ATTRIBUTES, bInitialOwner bool, lpName string) (win.HANDLE, error) {
	handle, err := win.CreateMutex(lpMutexAttributes, bInitialOwner, syscall.StringToUTF16Ptr(lpName))
	return handle, err
}

func CloseHandle(handle win.HANDLE) bool {
	return win.CloseHandle(handle)
}
