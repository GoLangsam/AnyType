// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type WriterAtChan interface { // bidirectional channel
	WriterAtROnlyChan // aka "<-chan" - receive only
	WriterAtSOnlyChan // aka "chan<-" - send only
}

type WriterAtROnlyChan interface { // receive-only channel
	RequestWriterAt() (dat io.WriterAt)        // the receive function - aka "some-new-WriterAt-var := <-MyKind"
	TryWriterAt() (dat io.WriterAt, open bool) // the multi-valued comma-ok receive function - aka "some-new-WriterAt-var, ok := <-MyKind"
}

type WriterAtSOnlyChan interface { // send-only channel
	ProvideWriterAt(dat io.WriterAt) // the send function - aka "MyKind <- some WriterAt"
}
