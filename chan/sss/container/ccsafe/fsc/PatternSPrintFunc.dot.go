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

// PatternSPrintFunc is a simple helper for PipePatternSFunc
func PatternSPrintFunc(prefix string) func(fp fs.PatternS) fs.PatternS {
	return func(fp fs.PatternS) fs.PatternS {
		fmt.Println(prefix, fp.String())
		return fp
	}
}
