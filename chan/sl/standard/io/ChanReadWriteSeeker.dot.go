// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type ReadWriteSeekerChan interface { // bidirectional channel
	ReadWriteSeekerROnlyChan // aka "<-chan" - receive only
	ReadWriteSeekerSOnlyChan // aka "chan<-" - send only
}

type ReadWriteSeekerROnlyChan interface { // receive-only channel
	RequestReadWriteSeeker() (dat io.ReadWriteSeeker)        // the receive function - aka "some-new-ReadWriteSeeker-var := <-MyKind"
	TryReadWriteSeeker() (dat io.ReadWriteSeeker, open bool) // the multi-valued comma-ok receive function - aka "some-new-ReadWriteSeeker-var, ok := <-MyKind"
}

type ReadWriteSeekerSOnlyChan interface { // send-only channel
	ProvideReadWriteSeeker(dat io.ReadWriteSeeker) // the send function - aka "MyKind <- some ReadWriteSeeker"
}
