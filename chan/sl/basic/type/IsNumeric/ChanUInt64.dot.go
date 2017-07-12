// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type UInt64Chan interface { // bidirectional channel
	UInt64ROnlyChan // aka "<-chan" - receive only
	UInt64SOnlyChan // aka "chan<-" - send only
}

type UInt64ROnlyChan interface { // receive-only channel
	RequestUInt64() (dat uint64)        // the receive function - aka "some-new-UInt64-var := <-MyKind"
	TryUInt64() (dat uint64, open bool) // the multi-valued comma-ok receive function - aka "some-new-UInt64-var, ok := <-MyKind"
}

type UInt64SOnlyChan interface { // send-only channel
	ProvideUInt64(dat uint64) // the send function - aka "MyKind <- some UInt64"
}
