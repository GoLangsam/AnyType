// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// ReadCloserChan represents a
// bidirectional
// channel
type ReadCloserChan interface {
	ReadCloserROnlyChan // aka "<-chan" - receive only
	ReadCloserSOnlyChan // aka "chan<-" - send only
}

// ReadCloserROnlyChan represents a
// receive-only
// channel
type ReadCloserROnlyChan interface {
	RequestReadCloser() (dat io.ReadCloser)        // the receive function - aka "MyReadCloser := <-MyReadCloserROnlyChan"
	TryReadCloser() (dat io.ReadCloser, open bool) // the multi-valued comma-ok receive function - aka "MyReadCloser, ok := <-MyReadCloserROnlyChan"
}

// ReadCloserSOnlyChan represents a
// send-only
// channel
type ReadCloserSOnlyChan interface {
	ProvideReadCloser(dat io.ReadCloser) // the send function - aka "MyKind <- some ReadCloser"
}
