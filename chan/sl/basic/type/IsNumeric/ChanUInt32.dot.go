// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// UInt32Chan represents a
// bidirectional
// channel
type UInt32Chan interface {
	UInt32ROnlyChan // aka "<-chan" - receive only
	UInt32SOnlyChan // aka "chan<-" - send only
}

// UInt32ROnlyChan represents a
// receive-only
// channel
type UInt32ROnlyChan interface {
	RequestUInt32() (dat uint32)        // the receive function - aka "MyUInt32 := <-MyUInt32ROnlyChan"
	TryUInt32() (dat uint32, open bool) // the multi-valued comma-ok receive function - aka "MyUInt32, ok := <-MyUInt32ROnlyChan"
}

// UInt32SOnlyChan represents a
// send-only
// channel
type UInt32SOnlyChan interface {
	ProvideUInt32(dat uint32) // the send function - aka "MyKind <- some UInt32"
}
