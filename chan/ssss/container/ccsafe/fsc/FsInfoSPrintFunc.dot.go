// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"fmt"

	"container/ccsafe/fs"
)

// FsInfoSPrintFunc is a simple helper for PipeFsInfoSFunc
func FsInfoSPrintFunc(prefix string) func(fp fs.FsInfoS) fs.FsInfoS {
	return func(fp fs.FsInfoS) fs.FsInfoS {
		fmt.Println(prefix, fp.String())
		return fp
	}
}
