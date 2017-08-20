// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// UInt16Chan represents a
// bidirectional
// channel
type UInt16Chan interface {
	UInt16ROnlyChan // aka "<-chan" - receive only
	UInt16SOnlyChan // aka "chan<-" - send only
}

// UInt16ROnlyChan represents a
// receive-only
// channel
type UInt16ROnlyChan interface {
	RequestUInt16() (dat uint16)        // the receive function - aka "MyUInt16 := <-MyUInt16ROnlyChan"
	TryUInt16() (dat uint16, open bool) // the multi-valued comma-ok receive function - aka "MyUInt16, ok := <-MyUInt16ROnlyChan"
}

// UInt16SOnlyChan represents a
// send-only
// channel
type UInt16SOnlyChan interface {
	ProvideUInt16(dat uint16) // the send function - aka "MyKind <- some UInt16"
}
