// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// PointerSChan represents a
// bidirectional
// channel
type PointerSChan interface {
	PointerSROnlyChan // aka "<-chan" - receive only
	PointerSSOnlyChan // aka "chan<-" - send only
}

// PointerSROnlyChan represents a
// receive-only
// channel
type PointerSROnlyChan interface {
	RequestPointerS() (dat []*SomeType)        // the receive function - aka "MyPointerS := <-MyPointerSROnlyChan"
	TryPointerS() (dat []*SomeType, open bool) // the multi-valued comma-ok receive function - aka "MyPointerS, ok := <-MyPointerSROnlyChan"
}

// PointerSSOnlyChan represents a
// send-only
// channel
type PointerSSOnlyChan interface {
	ProvidePointerS(dat []*SomeType) // the send function - aka "MyKind <- some PointerS"
}
