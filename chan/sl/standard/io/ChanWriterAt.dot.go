// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// WriterAtChan represents a
// bidirectional
// channel
type WriterAtChan interface {
	WriterAtROnlyChan // aka "<-chan" - receive only
	WriterAtSOnlyChan // aka "chan<-" - send only
}

// WriterAtROnlyChan represents a
// receive-only
// channel
type WriterAtROnlyChan interface {
	RequestWriterAt() (dat io.WriterAt)        // the receive function - aka "MyWriterAt := <-MyWriterAtROnlyChan"
	TryWriterAt() (dat io.WriterAt, open bool) // the multi-valued comma-ok receive function - aka "MyWriterAt, ok := <-MyWriterAtROnlyChan"
}

// WriterAtSOnlyChan represents a
// send-only
// channel
type WriterAtSOnlyChan interface {
	ProvideWriterAt(dat io.WriterAt) // the send function - aka "MyKind <- some WriterAt"
}
