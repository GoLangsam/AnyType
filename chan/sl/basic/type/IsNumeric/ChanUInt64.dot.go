// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// UInt64Chan represents a
// bidirectional
// channel
type UInt64Chan interface {
	UInt64ROnlyChan // aka "<-chan" - receive only
	UInt64SOnlyChan // aka "chan<-" - send only
}

// UInt64ROnlyChan represents a
// receive-only
// channel
type UInt64ROnlyChan interface {
	RequestUInt64() (dat uint64)        // the receive function - aka "MyUInt64 := <-MyUInt64ROnlyChan"
	TryUInt64() (dat uint64, open bool) // the multi-valued comma-ok receive function - aka "MyUInt64, ok := <-MyUInt64ROnlyChan"
}

// UInt64SOnlyChan represents a
// send-only
// channel
type UInt64SOnlyChan interface {
	ProvideUInt64(dat uint64) // the send function - aka "MyKind <- some UInt64"
}
