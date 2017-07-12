// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type PointerChan interface { // bidirectional channel
	PointerROnlyChan // aka "<-chan" - receive only
	PointerSOnlyChan // aka "chan<-" - send only
}

type PointerROnlyChan interface { // receive-only channel
	RequestPointer() (dat *SomeType)        // the receive function - aka "some-new-Pointer-var := <-MyKind"
	TryPointer() (dat *SomeType, open bool) // the multi-valued comma-ok receive function - aka "some-new-Pointer-var, ok := <-MyKind"
}

type PointerSOnlyChan interface { // send-only channel
	ProvidePointer(dat *SomeType) // the send function - aka "MyKind <- some Pointer"
}
