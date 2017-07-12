// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type SomeTypeChan interface { // bidirectional channel
	SomeTypeROnlyChan // aka "<-chan" - receive only
	SomeTypeSOnlyChan // aka "chan<-" - send only
}

type SomeTypeROnlyChan interface { // receive-only channel
	RequestSomeType() (dat SomeType)        // the receive function - aka "some-new-SomeType-var := <-MyKind"
	TrySomeType() (dat SomeType, open bool) // the multi-valued comma-ok receive function - aka "some-new-SomeType-var, ok := <-MyKind"
}

type SomeTypeSOnlyChan interface { // send-only channel
	ProvideSomeType(dat SomeType) // the send function - aka "MyKind <- some SomeType"
}