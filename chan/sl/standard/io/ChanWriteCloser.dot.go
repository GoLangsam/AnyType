// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type WriteCloserChan interface { // bidirectional channel
	WriteCloserROnlyChan // aka "<-chan" - receive only
	WriteCloserSOnlyChan // aka "chan<-" - send only
}

type WriteCloserROnlyChan interface { // receive-only channel
	RequestWriteCloser() (dat io.WriteCloser)        // the receive function - aka "some-new-WriteCloser-var := <-MyKind"
	TryWriteCloser() (dat io.WriteCloser, open bool) // the multi-valued comma-ok receive function - aka "some-new-WriteCloser-var, ok := <-MyKind"
}

type WriteCloserSOnlyChan interface { // send-only channel
	ProvideWriteCloser(dat io.WriteCloser) // the send function - aka "MyKind <- some WriteCloser"
}
