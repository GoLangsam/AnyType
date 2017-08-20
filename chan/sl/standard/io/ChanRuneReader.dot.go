// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// RuneReaderChan represents a
// bidirectional
// channel
type RuneReaderChan interface {
	RuneReaderROnlyChan // aka "<-chan" - receive only
	RuneReaderSOnlyChan // aka "chan<-" - send only
}

// RuneReaderROnlyChan represents a
// receive-only
// channel
type RuneReaderROnlyChan interface {
	RequestRuneReader() (dat io.RuneReader)        // the receive function - aka "MyRuneReader := <-MyRuneReaderROnlyChan"
	TryRuneReader() (dat io.RuneReader, open bool) // the multi-valued comma-ok receive function - aka "MyRuneReader, ok := <-MyRuneReaderROnlyChan"
}

// RuneReaderSOnlyChan represents a
// send-only
// channel
type RuneReaderSOnlyChan interface {
	ProvideRuneReader(dat io.RuneReader) // the send function - aka "MyKind <- some RuneReader"
}
