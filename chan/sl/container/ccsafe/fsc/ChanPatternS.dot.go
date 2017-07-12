// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/fs"
)

type PatternSChan interface { // bidirectional channel
	PatternSROnlyChan // aka "<-chan" - receive only
	PatternSSOnlyChan // aka "chan<-" - send only
}

type PatternSROnlyChan interface { // receive-only channel
	RequestPatternS() (dat fs.PatternS)        // the receive function - aka "some-new-PatternS-var := <-MyKind"
	TryPatternS() (dat fs.PatternS, open bool) // the multi-valued comma-ok receive function - aka "some-new-PatternS-var, ok := <-MyKind"
}

type PatternSSOnlyChan interface { // send-only channel
	ProvidePatternS(dat fs.PatternS) // the send function - aka "MyKind <- some PatternS"
}
