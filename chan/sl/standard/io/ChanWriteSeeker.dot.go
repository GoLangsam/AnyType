// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// WriteSeekerChan represents a
// bidirectional
// channel
type WriteSeekerChan interface {
	WriteSeekerROnlyChan // aka "<-chan" - receive only
	WriteSeekerSOnlyChan // aka "chan<-" - send only
}

// WriteSeekerROnlyChan represents a
// receive-only
// channel
type WriteSeekerROnlyChan interface {
	RequestWriteSeeker() (dat io.WriteSeeker)        // the receive function - aka "MyWriteSeeker := <-MyWriteSeekerROnlyChan"
	TryWriteSeeker() (dat io.WriteSeeker, open bool) // the multi-valued comma-ok receive function - aka "MyWriteSeeker, ok := <-MyWriteSeekerROnlyChan"
}

// WriteSeekerSOnlyChan represents a
// send-only
// channel
type WriteSeekerSOnlyChan interface {
	ProvideWriteSeeker(dat io.WriteSeeker) // the send function - aka "MyKind <- some WriteSeeker"
}
