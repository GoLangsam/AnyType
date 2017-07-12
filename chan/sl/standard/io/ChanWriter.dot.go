// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type WriterChan interface { // bidirectional channel
	WriterROnlyChan // aka "<-chan" - receive only
	WriterSOnlyChan // aka "chan<-" - send only
}

type WriterROnlyChan interface { // receive-only channel
	RequestWriter() (dat io.Writer)        // the receive function - aka "some-new-Writer-var := <-MyKind"
	TryWriter() (dat io.Writer, open bool) // the multi-valued comma-ok receive function - aka "some-new-Writer-var, ok := <-MyKind"
}

type WriterSOnlyChan interface { // send-only channel
	ProvideWriter(dat io.Writer) // the send function - aka "MyKind <- some Writer"
}
