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

// PatternPrintFunc is a simple helper for PipePatternFunc
func PatternPrintFunc(prefix string) func(fp *fs.Pattern) *fs.Pattern {
	return func(fp *fs.Pattern) *fs.Pattern {
		fmt.Println(prefix, fp.String())
		return fp
	}
}
