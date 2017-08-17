// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type ReadWriterChan interface { // bidirectional channel
	ReadWriterROnlyChan // aka "<-chan" - receive only
	ReadWriterSOnlyChan // aka "chan<-" - send only
}

type ReadWriterROnlyChan interface { // receive-only channel
	RequestReadWriter() (dat io.ReadWriter)        // the receive function - aka "some-new-ReadWriter-var := <-MyKind"
	TryReadWriter() (dat io.ReadWriter, open bool) // the multi-valued comma-ok receive function - aka "some-new-ReadWriter-var, ok := <-MyKind"
}

type ReadWriterSOnlyChan interface { // send-only channel
	ProvideReadWriter(dat io.ReadWriter) // the send function - aka "MyKind <- some ReadWriter"
}
