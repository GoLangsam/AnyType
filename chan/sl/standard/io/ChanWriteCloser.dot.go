// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// WriteCloserChan represents a
// bidirectional
// channel
type WriteCloserChan interface {
	WriteCloserROnlyChan // aka "<-chan" - receive only
	WriteCloserSOnlyChan // aka "chan<-" - send only
}

// WriteCloserROnlyChan represents a
// receive-only
// channel
type WriteCloserROnlyChan interface {
	RequestWriteCloser() (dat io.WriteCloser)        // the receive function - aka "MyWriteCloser := <-MyWriteCloserROnlyChan"
	TryWriteCloser() (dat io.WriteCloser, open bool) // the multi-valued comma-ok receive function - aka "MyWriteCloser, ok := <-MyWriteCloserROnlyChan"
}

// WriteCloserSOnlyChan represents a
// send-only
// channel
type WriteCloserSOnlyChan interface {
	ProvideWriteCloser(dat io.WriteCloser) // the send function - aka "MyKind <- some WriteCloser"
}
