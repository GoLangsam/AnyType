// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/dot"
)

// DotChan represents a
// bidirectional
// channel
type DotChan interface {
	DotROnlyChan // aka "<-chan" - receive only
	DotSOnlyChan // aka "chan<-" - send only
}

// DotROnlyChan represents a
// receive-only
// channel
type DotROnlyChan interface {
	RequestDot() (dat dot.Dot)        // the receive function - aka "MyDot := <-MyDotROnlyChan"
	TryDot() (dat dot.Dot, open bool) // the multi-valued comma-ok receive function - aka "MyDot, ok := <-MyDotROnlyChan"
}

// DotSOnlyChan represents a
// send-only
// channel
type DotSOnlyChan interface {
	ProvideDot(dat dot.Dot) // the send function - aka "MyKind <- some Dot"
}
