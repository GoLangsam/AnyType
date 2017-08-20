// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// ReaderFromChan represents a
// bidirectional
// channel
type ReaderFromChan interface {
	ReaderFromROnlyChan // aka "<-chan" - receive only
	ReaderFromSOnlyChan // aka "chan<-" - send only
}

// ReaderFromROnlyChan represents a
// receive-only
// channel
type ReaderFromROnlyChan interface {
	RequestReaderFrom() (dat io.ReaderFrom)        // the receive function - aka "MyReaderFrom := <-MyReaderFromROnlyChan"
	TryReaderFrom() (dat io.ReaderFrom, open bool) // the multi-valued comma-ok receive function - aka "MyReaderFrom, ok := <-MyReaderFromROnlyChan"
}

// ReaderFromSOnlyChan represents a
// send-only
// channel
type ReaderFromSOnlyChan interface {
	ProvideReaderFrom(dat io.ReaderFrom) // the send function - aka "MyKind <- some ReaderFrom"
}
