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

// FsFoldSPrintFunc is a simple helper for PipeFsFoldSFunc
func FsFoldSPrintFunc(prefix string) func(fp fs.FsFoldS) fs.FsFoldS {
	return func(fp fs.FsFoldS) fs.FsFoldS {
		fmt.Println(prefix, fp.String())
		return fp
	}
}
