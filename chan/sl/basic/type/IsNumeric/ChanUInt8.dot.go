// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// UInt8Chan represents a
// bidirectional
// channel
type UInt8Chan interface {
	UInt8ROnlyChan // aka "<-chan" - receive only
	UInt8SOnlyChan // aka "chan<-" - send only
}

// UInt8ROnlyChan represents a
// receive-only
// channel
type UInt8ROnlyChan interface {
	RequestUInt8() (dat uint8)        // the receive function - aka "MyUInt8 := <-MyUInt8ROnlyChan"
	TryUInt8() (dat uint8, open bool) // the multi-valued comma-ok receive function - aka "MyUInt8, ok := <-MyUInt8ROnlyChan"
}

// UInt8SOnlyChan represents a
// send-only
// channel
type UInt8SOnlyChan interface {
	ProvideUInt8(dat uint8) // the send function - aka "MyKind <- some UInt8"
}
