// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// PointerChan represents a
// bidirectional
// channel
type PointerChan interface {
	PointerROnlyChan // aka "<-chan" - receive only
	PointerSOnlyChan // aka "chan<-" - send only
}

// PointerROnlyChan represents a
// receive-only
// channel
type PointerROnlyChan interface {
	RequestPointer() (dat *SomeType)        // the receive function - aka "MyPointer := <-MyPointerROnlyChan"
	TryPointer() (dat *SomeType, open bool) // the multi-valued comma-ok receive function - aka "MyPointer, ok := <-MyPointerROnlyChan"
}

// PointerSOnlyChan represents a
// send-only
// channel
type PointerSOnlyChan interface {
	ProvidePointer(dat *SomeType) // the send function - aka "MyKind <- some Pointer"
}
