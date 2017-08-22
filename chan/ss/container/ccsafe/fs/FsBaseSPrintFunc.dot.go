// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"fmt"

	"github.com/golangsam/container/ccsafe/fs"
)

// FsBaseSPrintFunc is a simple helper for PipeFsBaseSFunc
func FsBaseSPrintFunc(prefix string) func(fp fs.FsBaseS) fs.FsBaseS {
	return func(fp fs.FsBaseS) fs.FsBaseS {
		fmt.Println(prefix, fp.String())
		return fp
	}
}
