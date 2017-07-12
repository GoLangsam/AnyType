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

// FsFoldPrintFunc is a simple helper for PipeFsFoldFunc
func FsFoldPrintFunc(prefix string) func(fp *fs.FsFold) *fs.FsFold {
	return func(fp *fs.FsFold) *fs.FsFold {
		fmt.Println(prefix, fp.String())
		return fp
	}
}
