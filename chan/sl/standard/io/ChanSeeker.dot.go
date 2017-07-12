// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type SeekerChan interface { // bidirectional channel
	SeekerROnlyChan // aka "<-chan" - receive only
	SeekerSOnlyChan // aka "chan<-" - send only
}

type SeekerROnlyChan interface { // receive-only channel
	RequestSeeker() (dat io.Seeker)        // the receive function - aka "some-new-Seeker-var := <-MyKind"
	TrySeeker() (dat io.Seeker, open bool) // the multi-valued comma-ok receive function - aka "some-new-Seeker-var, ok := <-MyKind"
}

type SeekerSOnlyChan interface { // send-only channel
	ProvideSeeker(dat io.Seeker) // the send function - aka "MyKind <- some Seeker"
}
