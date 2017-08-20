// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// WriterToChan represents a
// bidirectional
// channel
type WriterToChan interface {
	WriterToROnlyChan // aka "<-chan" - receive only
	WriterToSOnlyChan // aka "chan<-" - send only
}

// WriterToROnlyChan represents a
// receive-only
// channel
type WriterToROnlyChan interface {
	RequestWriterTo() (dat io.WriterTo)        // the receive function - aka "MyWriterTo := <-MyWriterToROnlyChan"
	TryWriterTo() (dat io.WriterTo, open bool) // the multi-valued comma-ok receive function - aka "MyWriterTo, ok := <-MyWriterToROnlyChan"
}

// WriterToSOnlyChan represents a
// send-only
// channel
type WriterToSOnlyChan interface {
	ProvideWriterTo(dat io.WriterTo) // the send function - aka "MyKind <- some WriterTo"
}
