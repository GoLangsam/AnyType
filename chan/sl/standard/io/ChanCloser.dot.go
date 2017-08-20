// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// CloserChan represents a
// bidirectional
// channel
type CloserChan interface {
	CloserROnlyChan // aka "<-chan" - receive only
	CloserSOnlyChan // aka "chan<-" - send only
}

// CloserROnlyChan represents a
// receive-only
// channel
type CloserROnlyChan interface {
	RequestCloser() (dat io.Closer)        // the receive function - aka "MyCloser := <-MyCloserROnlyChan"
	TryCloser() (dat io.Closer, open bool) // the multi-valued comma-ok receive function - aka "MyCloser, ok := <-MyCloserROnlyChan"
}

// CloserSOnlyChan represents a
// send-only
// channel
type CloserSOnlyChan interface {
	ProvideCloser(dat io.Closer) // the send function - aka "MyKind <- some Closer"
}
