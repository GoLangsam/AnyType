// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// ReaderAtChan represents a
// bidirectional
// channel
type ReaderAtChan interface {
	ReaderAtROnlyChan // aka "<-chan" - receive only
	ReaderAtSOnlyChan // aka "chan<-" - send only
}

// ReaderAtROnlyChan represents a
// receive-only
// channel
type ReaderAtROnlyChan interface {
	RequestReaderAt() (dat io.ReaderAt)        // the receive function - aka "MyReaderAt := <-MyReaderAtROnlyChan"
	TryReaderAt() (dat io.ReaderAt, open bool) // the multi-valued comma-ok receive function - aka "MyReaderAt, ok := <-MyReaderAtROnlyChan"
}

// ReaderAtSOnlyChan represents a
// send-only
// channel
type ReaderAtSOnlyChan interface {
	ProvideReaderAt(dat io.ReaderAt) // the send function - aka "MyKind <- some ReaderAt"
}
