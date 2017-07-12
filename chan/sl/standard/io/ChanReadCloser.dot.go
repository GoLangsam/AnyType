// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type ReadCloserChan interface { // bidirectional channel
	ReadCloserROnlyChan // aka "<-chan" - receive only
	ReadCloserSOnlyChan // aka "chan<-" - send only
}

type ReadCloserROnlyChan interface { // receive-only channel
	RequestReadCloser() (dat io.ReadCloser)        // the receive function - aka "some-new-ReadCloser-var := <-MyKind"
	TryReadCloser() (dat io.ReadCloser, open bool) // the multi-valued comma-ok receive function - aka "some-new-ReadCloser-var, ok := <-MyKind"
}

type ReadCloserSOnlyChan interface { // send-only channel
	ProvideReadCloser(dat io.ReadCloser) // the send function - aka "MyKind <- some ReadCloser"
}