// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// ByteReaderChan represents a
// bidirectional
// channel
type ByteReaderChan interface {
	ByteReaderROnlyChan // aka "<-chan" - receive only
	ByteReaderSOnlyChan // aka "chan<-" - send only
}

// ByteReaderROnlyChan represents a
// receive-only
// channel
type ByteReaderROnlyChan interface {
	RequestByteReader() (dat io.ByteReader)        // the receive function - aka "MyByteReader := <-MyByteReaderROnlyChan"
	TryByteReader() (dat io.ByteReader, open bool) // the multi-valued comma-ok receive function - aka "MyByteReader, ok := <-MyByteReaderROnlyChan"
}

// ByteReaderSOnlyChan represents a
// send-only
// channel
type ByteReaderSOnlyChan interface {
	ProvideByteReader(dat io.ByteReader) // the send function - aka "MyKind <- some ByteReader"
}
