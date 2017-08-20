// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// IntChan represents a
// bidirectional
// channel
type IntChan interface {
	IntROnlyChan // aka "<-chan" - receive only
	IntSOnlyChan // aka "chan<-" - send only
}

// IntROnlyChan represents a
// receive-only
// channel
type IntROnlyChan interface {
	RequestInt() (dat int)        // the receive function - aka "MyInt := <-MyIntROnlyChan"
	TryInt() (dat int, open bool) // the multi-valued comma-ok receive function - aka "MyInt, ok := <-MyIntROnlyChan"
}

// IntSOnlyChan represents a
// send-only
// channel
type IntSOnlyChan interface {
	ProvideInt(dat int) // the send function - aka "MyKind <- some Int"
}
