// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// ReadSeekerChan represents a
// bidirectional
// channel
type ReadSeekerChan interface {
	ReadSeekerROnlyChan // aka "<-chan" - receive only
	ReadSeekerSOnlyChan // aka "chan<-" - send only
}

// ReadSeekerROnlyChan represents a
// receive-only
// channel
type ReadSeekerROnlyChan interface {
	RequestReadSeeker() (dat io.ReadSeeker)        // the receive function - aka "MyReadSeeker := <-MyReadSeekerROnlyChan"
	TryReadSeeker() (dat io.ReadSeeker, open bool) // the multi-valued comma-ok receive function - aka "MyReadSeeker, ok := <-MyReadSeekerROnlyChan"
}

// ReadSeekerSOnlyChan represents a
// send-only
// channel
type ReadSeekerSOnlyChan interface {
	ProvideReadSeeker(dat io.ReadSeeker) // the send function - aka "MyKind <- some ReadSeeker"
}
