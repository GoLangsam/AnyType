// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type ByteReaderChan interface { // bidirectional channel
	ByteReaderROnlyChan // aka "<-chan" - receive only
	ByteReaderSOnlyChan // aka "chan<-" - send only
}

type ByteReaderROnlyChan interface { // receive-only channel
	RequestByteReader() (dat io.ByteReader)        // the receive function - aka "some-new-ByteReader-var := <-MyKind"
	TryByteReader() (dat io.ByteReader, open bool) // the multi-valued comma-ok receive function - aka "some-new-ByteReader-var, ok := <-MyKind"
}

type ByteReaderSOnlyChan interface { // send-only channel
	ProvideByteReader(dat io.ByteReader) // the send function - aka "MyKind <- some ByteReader"
}
