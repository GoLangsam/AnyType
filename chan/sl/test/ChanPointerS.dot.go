// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type PointerSChan interface { // bidirectional channel
	PointerSROnlyChan // aka "<-chan" - receive only
	PointerSSOnlyChan // aka "chan<-" - send only
}

type PointerSROnlyChan interface { // receive-only channel
	RequestPointerS() (dat []*SomeType)        // the receive function - aka "some-new-PointerS-var := <-MyKind"
	TryPointerS() (dat []*SomeType, open bool) // the multi-valued comma-ok receive function - aka "some-new-PointerS-var, ok := <-MyKind"
}

type PointerSSOnlyChan interface { // send-only channel
	ProvidePointerS(dat []*SomeType) // the send function - aka "MyKind <- some PointerS"
}
