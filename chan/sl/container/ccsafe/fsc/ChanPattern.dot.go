// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// PatternChan represents a
// bidirectional
// channel
type PatternChan interface {
	PatternROnlyChan // aka "<-chan" - receive only
	PatternSOnlyChan // aka "chan<-" - send only
}

// PatternROnlyChan represents a
// receive-only
// channel
type PatternROnlyChan interface {
	RequestPattern() (dat *fs.Pattern)        // the receive function - aka "MyPattern := <-MyPatternROnlyChan"
	TryPattern() (dat *fs.Pattern, open bool) // the multi-valued comma-ok receive function - aka "MyPattern, ok := <-MyPatternROnlyChan"
}

// PatternSOnlyChan represents a
// send-only
// channel
type PatternSOnlyChan interface {
	ProvidePattern(dat *fs.Pattern) // the send function - aka "MyKind <- some Pattern"
}
