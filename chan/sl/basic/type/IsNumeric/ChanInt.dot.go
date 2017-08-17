// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type IntChan interface { // bidirectional channel
	IntROnlyChan // aka "<-chan" - receive only
	IntSOnlyChan // aka "chan<-" - send only
}

type IntROnlyChan interface { // receive-only channel
	RequestInt() (dat int)        // the receive function - aka "some-new-Int-var := <-MyKind"
	TryInt() (dat int, open bool) // the multi-valued comma-ok receive function - aka "some-new-Int-var, ok := <-MyKind"
}

type IntSOnlyChan interface { // send-only channel
	ProvideInt(dat int) // the send function - aka "MyKind <- some Int"
}
