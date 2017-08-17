// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type WriterToChan interface { // bidirectional channel
	WriterToROnlyChan // aka "<-chan" - receive only
	WriterToSOnlyChan // aka "chan<-" - send only
}

type WriterToROnlyChan interface { // receive-only channel
	RequestWriterTo() (dat io.WriterTo)        // the receive function - aka "some-new-WriterTo-var := <-MyKind"
	TryWriterTo() (dat io.WriterTo, open bool) // the multi-valued comma-ok receive function - aka "some-new-WriterTo-var, ok := <-MyKind"
}

type WriterToSOnlyChan interface { // send-only channel
	ProvideWriterTo(dat io.WriterTo) // the send function - aka "MyKind <- some WriterTo"
}
