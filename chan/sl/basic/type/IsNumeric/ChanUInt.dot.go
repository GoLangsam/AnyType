// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type UIntChan interface { // bidirectional channel
	UIntROnlyChan // aka "<-chan" - receive only
	UIntSOnlyChan // aka "chan<-" - send only
}

type UIntROnlyChan interface { // receive-only channel
	RequestUInt() (dat uint)        // the receive function - aka "some-new-UInt-var := <-MyKind"
	TryUInt() (dat uint, open bool) // the multi-valued comma-ok receive function - aka "some-new-UInt-var, ok := <-MyKind"
}

type UIntSOnlyChan interface { // send-only channel
	ProvideUInt(dat uint) // the send function - aka "MyKind <- some UInt"
}
