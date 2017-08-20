// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/list"
)

// ElementSChan represents a
// bidirectional
// channel
type ElementSChan interface {
	ElementSROnlyChan // aka "<-chan" - receive only
	ElementSSOnlyChan // aka "chan<-" - send only
}

// ElementSROnlyChan represents a
// receive-only
// channel
type ElementSROnlyChan interface {
	RequestElementS() (dat []list.Element)        // the receive function - aka "MyElementS := <-MyElementSROnlyChan"
	TryElementS() (dat []list.Element, open bool) // the multi-valued comma-ok receive function - aka "MyElementS, ok := <-MyElementSROnlyChan"
}

// ElementSSOnlyChan represents a
// send-only
// channel
type ElementSSOnlyChan interface {
	ProvideElementS(dat []list.Element) // the send function - aka "MyKind <- some ElementS"
}
