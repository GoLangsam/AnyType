// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bytes"
)

// BufferChan represents a
// bidirectional
// channel
type BufferChan interface {
	BufferROnlyChan // aka "<-chan" - receive only
	BufferSOnlyChan // aka "chan<-" - send only
}

// BufferROnlyChan represents a
// receive-only
// channel
type BufferROnlyChan interface {
	RequestBuffer() (dat bytes.Buffer)        // the receive function - aka "MyBuffer := <-MyBufferROnlyChan"
	TryBuffer() (dat bytes.Buffer, open bool) // the multi-valued comma-ok receive function - aka "MyBuffer, ok := <-MyBufferROnlyChan"
}

// BufferSOnlyChan represents a
// send-only
// channel
type BufferSOnlyChan interface {
	ProvideBuffer(dat bytes.Buffer) // the send function - aka "MyKind <- some Buffer"
}
