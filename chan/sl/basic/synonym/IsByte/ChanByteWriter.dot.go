// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type ByteWriterChan interface { // bidirectional channel
	ByteWriterROnlyChan // aka "<-chan" - receive only
	ByteWriterSOnlyChan // aka "chan<-" - send only
}

type ByteWriterROnlyChan interface { // receive-only channel
	RequestByteWriter() (dat io.ByteWriter)        // the receive function - aka "some-new-ByteWriter-var := <-MyKind"
	TryByteWriter() (dat io.ByteWriter, open bool) // the multi-valued comma-ok receive function - aka "some-new-ByteWriter-var, ok := <-MyKind"
}

type ByteWriterSOnlyChan interface { // send-only channel
	ProvideByteWriter(dat io.ByteWriter) // the send function - aka "MyKind <- some ByteWriter"
}
