// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type UInt32Chan interface { // bidirectional channel
	UInt32ROnlyChan // aka "<-chan" - receive only
	UInt32SOnlyChan // aka "chan<-" - send only
}

type UInt32ROnlyChan interface { // receive-only channel
	RequestUInt32() (dat uint32)        // the receive function - aka "some-new-UInt32-var := <-MyKind"
	TryUInt32() (dat uint32, open bool) // the multi-valued comma-ok receive function - aka "some-new-UInt32-var, ok := <-MyKind"
}

type UInt32SOnlyChan interface { // send-only channel
	ProvideUInt32(dat uint32) // the send function - aka "MyKind <- some UInt32"
}
