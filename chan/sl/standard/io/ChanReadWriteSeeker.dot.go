// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// ReadWriteSeekerChan represents a
// bidirectional
// channel
type ReadWriteSeekerChan interface {
	ReadWriteSeekerROnlyChan // aka "<-chan" - receive only
	ReadWriteSeekerSOnlyChan // aka "chan<-" - send only
}

// ReadWriteSeekerROnlyChan represents a
// receive-only
// channel
type ReadWriteSeekerROnlyChan interface {
	RequestReadWriteSeeker() (dat io.ReadWriteSeeker)        // the receive function - aka "MyReadWriteSeeker := <-MyReadWriteSeekerROnlyChan"
	TryReadWriteSeeker() (dat io.ReadWriteSeeker, open bool) // the multi-valued comma-ok receive function - aka "MyReadWriteSeeker, ok := <-MyReadWriteSeekerROnlyChan"
}

// ReadWriteSeekerSOnlyChan represents a
// send-only
// channel
type ReadWriteSeekerSOnlyChan interface {
	ProvideReadWriteSeeker(dat io.ReadWriteSeeker) // the send function - aka "MyKind <- some ReadWriteSeeker"
}
