// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type CloserChan interface { // bidirectional channel
	CloserROnlyChan // aka "<-chan" - receive only
	CloserSOnlyChan // aka "chan<-" - send only
}

type CloserROnlyChan interface { // receive-only channel
	RequestCloser() (dat io.Closer)        // the receive function - aka "some-new-Closer-var := <-MyKind"
	TryCloser() (dat io.Closer, open bool) // the multi-valued comma-ok receive function - aka "some-new-Closer-var, ok := <-MyKind"
}

type CloserSOnlyChan interface { // send-only channel
	ProvideCloser(dat io.Closer) // the send function - aka "MyKind <- some Closer"
}
