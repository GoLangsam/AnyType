// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// ByteWriterChan represents a
// bidirectional
// channel
type ByteWriterChan interface {
	ByteWriterROnlyChan // aka "<-chan" - receive only
	ByteWriterSOnlyChan // aka "chan<-" - send only
}

// ByteWriterROnlyChan represents a
// receive-only
// channel
type ByteWriterROnlyChan interface {
	RequestByteWriter() (dat io.ByteWriter)        // the receive function - aka "MyByteWriter := <-MyByteWriterROnlyChan"
	TryByteWriter() (dat io.ByteWriter, open bool) // the multi-valued comma-ok receive function - aka "MyByteWriter, ok := <-MyByteWriterROnlyChan"
}

// ByteWriterSOnlyChan represents a
// send-only
// channel
type ByteWriterSOnlyChan interface {
	ProvideByteWriter(dat io.ByteWriter) // the send function - aka "MyKind <- some ByteWriter"
}
