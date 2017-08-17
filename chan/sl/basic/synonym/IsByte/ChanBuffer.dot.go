// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bytes"
)

type BufferChan interface { // bidirectional channel
	BufferROnlyChan // aka "<-chan" - receive only
	BufferSOnlyChan // aka "chan<-" - send only
}

type BufferROnlyChan interface { // receive-only channel
	RequestBuffer() (dat bytes.Buffer)        // the receive function - aka "some-new-Buffer-var := <-MyKind"
	TryBuffer() (dat bytes.Buffer, open bool) // the multi-valued comma-ok receive function - aka "some-new-Buffer-var, ok := <-MyKind"
}

type BufferSOnlyChan interface { // send-only channel
	ProvideBuffer(dat bytes.Buffer) // the send function - aka "MyKind <- some Buffer"
}
