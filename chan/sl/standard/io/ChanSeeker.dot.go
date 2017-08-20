// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// SeekerChan represents a
// bidirectional
// channel
type SeekerChan interface {
	SeekerROnlyChan // aka "<-chan" - receive only
	SeekerSOnlyChan // aka "chan<-" - send only
}

// SeekerROnlyChan represents a
// receive-only
// channel
type SeekerROnlyChan interface {
	RequestSeeker() (dat io.Seeker)        // the receive function - aka "MySeeker := <-MySeekerROnlyChan"
	TrySeeker() (dat io.Seeker, open bool) // the multi-valued comma-ok receive function - aka "MySeeker, ok := <-MySeekerROnlyChan"
}

// SeekerSOnlyChan represents a
// send-only
// channel
type SeekerSOnlyChan interface {
	ProvideSeeker(dat io.Seeker) // the send function - aka "MyKind <- some Seeker"
}
