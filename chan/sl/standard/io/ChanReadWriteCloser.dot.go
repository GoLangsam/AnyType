// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type ReadWriteCloserChan interface { // bidirectional channel
	ReadWriteCloserROnlyChan // aka "<-chan" - receive only
	ReadWriteCloserSOnlyChan // aka "chan<-" - send only
}

type ReadWriteCloserROnlyChan interface { // receive-only channel
	RequestReadWriteCloser() (dat io.ReadWriteCloser)        // the receive function - aka "some-new-ReadWriteCloser-var := <-MyKind"
	TryReadWriteCloser() (dat io.ReadWriteCloser, open bool) // the multi-valued comma-ok receive function - aka "some-new-ReadWriteCloser-var, ok := <-MyKind"
}

type ReadWriteCloserSOnlyChan interface { // send-only channel
	ProvideReadWriteCloser(dat io.ReadWriteCloser) // the send function - aka "MyKind <- some ReadWriteCloser"
}
