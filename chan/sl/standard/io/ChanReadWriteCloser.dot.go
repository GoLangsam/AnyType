// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// ReadWriteCloserChan represents a
// bidirectional
// channel
type ReadWriteCloserChan interface {
	ReadWriteCloserROnlyChan // aka "<-chan" - receive only
	ReadWriteCloserSOnlyChan // aka "chan<-" - send only
}

// ReadWriteCloserROnlyChan represents a
// receive-only
// channel
type ReadWriteCloserROnlyChan interface {
	RequestReadWriteCloser() (dat io.ReadWriteCloser)        // the receive function - aka "MyReadWriteCloser := <-MyReadWriteCloserROnlyChan"
	TryReadWriteCloser() (dat io.ReadWriteCloser, open bool) // the multi-valued comma-ok receive function - aka "MyReadWriteCloser, ok := <-MyReadWriteCloserROnlyChan"
}

// ReadWriteCloserSOnlyChan represents a
// send-only
// channel
type ReadWriteCloserSOnlyChan interface {
	ProvideReadWriteCloser(dat io.ReadWriteCloser) // the send function - aka "MyKind <- some ReadWriteCloser"
}
