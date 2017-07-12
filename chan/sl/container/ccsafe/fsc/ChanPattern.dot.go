// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/fs"
)

type PatternChan interface { // bidirectional channel
	PatternROnlyChan // aka "<-chan" - receive only
	PatternSOnlyChan // aka "chan<-" - send only
}

type PatternROnlyChan interface { // receive-only channel
	RequestPattern() (dat *fs.Pattern)        // the receive function - aka "some-new-Pattern-var := <-MyKind"
	TryPattern() (dat *fs.Pattern, open bool) // the multi-valued comma-ok receive function - aka "some-new-Pattern-var, ok := <-MyKind"
}

type PatternSOnlyChan interface { // send-only channel
	ProvidePattern(dat *fs.Pattern) // the send function - aka "MyKind <- some Pattern"
}
