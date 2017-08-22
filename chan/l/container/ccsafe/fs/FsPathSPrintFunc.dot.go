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

// FsPathSPrintFunc is a simple helper for PipeFsPathSFunc
func FsPathSPrintFunc(prefix string) func(fp fs.FsPathS) fs.FsPathS {
	return func(fp fs.FsPathS) fs.FsPathS {
		fmt.Println(prefix, fp.String())
		return fp
	}
}
