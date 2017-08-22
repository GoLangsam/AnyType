// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"fmt"

	"github.com/golangsam/container/ccsafe/fs"
)

// FsFileSPrintFunc is a simple helper for PipeFsFileSFunc
func FsFileSPrintFunc(prefix string) func(fp fs.FsFileS) fs.FsFileS {
	return func(fp fs.FsFileS) fs.FsFileS {
		fmt.Println(prefix, fp.String())
		return fp
	}
}
