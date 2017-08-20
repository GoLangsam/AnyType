// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// UIntChan represents a
// bidirectional
// channel
type UIntChan interface {
	UIntROnlyChan // aka "<-chan" - receive only
	UIntSOnlyChan // aka "chan<-" - send only
}

// UIntROnlyChan represents a
// receive-only
// channel
type UIntROnlyChan interface {
	RequestUInt() (dat uint)        // the receive function - aka "MyUInt := <-MyUIntROnlyChan"
	TryUInt() (dat uint, open bool) // the multi-valued comma-ok receive function - aka "MyUInt, ok := <-MyUIntROnlyChan"
}

// UIntSOnlyChan represents a
// send-only
// channel
type UIntSOnlyChan interface {
	ProvideUInt(dat uint) // the send function - aka "MyKind <- some UInt"
}
