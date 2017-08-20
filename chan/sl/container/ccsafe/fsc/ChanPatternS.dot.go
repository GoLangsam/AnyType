// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// PatternSChan represents a
// bidirectional
// channel
type PatternSChan interface {
	PatternSROnlyChan // aka "<-chan" - receive only
	PatternSSOnlyChan // aka "chan<-" - send only
}

// PatternSROnlyChan represents a
// receive-only
// channel
type PatternSROnlyChan interface {
	RequestPatternS() (dat fs.PatternS)        // the receive function - aka "MyPatternS := <-MyPatternSROnlyChan"
	TryPatternS() (dat fs.PatternS, open bool) // the multi-valued comma-ok receive function - aka "MyPatternS, ok := <-MyPatternSROnlyChan"
}

// PatternSSOnlyChan represents a
// send-only
// channel
type PatternSSOnlyChan interface {
	ProvidePatternS(dat fs.PatternS) // the send function - aka "MyKind <- some PatternS"
}
