// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package dot

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/dot"
)

type DotChan interface { // bidirectional channel
	DotROnlyChan // aka "<-chan" - receive only
	DotSOnlyChan // aka "chan<-" - send only
}

type DotROnlyChan interface { // receive-only channel
	RequestDot() (dat dot.Dot)        // the receive function - aka "some-new-Dot-var := <-MyKind"
	TryDot() (dat dot.Dot, open bool) // the multi-valued comma-ok receive function - aka "some-new-Dot-var, ok := <-MyKind"
}

type DotSOnlyChan interface { // send-only channel
	ProvideDot(dat dot.Dot) // the send function - aka "MyKind <- some Dot"
}
