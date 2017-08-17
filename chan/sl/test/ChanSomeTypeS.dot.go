// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type SomeTypeSChan interface { // bidirectional channel
	SomeTypeSROnlyChan // aka "<-chan" - receive only
	SomeTypeSSOnlyChan // aka "chan<-" - send only
}

type SomeTypeSROnlyChan interface { // receive-only channel
	RequestSomeTypeS() (dat []SomeType)        // the receive function - aka "some-new-SomeTypeS-var := <-MyKind"
	TrySomeTypeS() (dat []SomeType, open bool) // the multi-valued comma-ok receive function - aka "some-new-SomeTypeS-var, ok := <-MyKind"
}

type SomeTypeSSOnlyChan interface { // send-only channel
	ProvideSomeTypeS(dat []SomeType) // the send function - aka "MyKind <- some SomeTypeS"
}
