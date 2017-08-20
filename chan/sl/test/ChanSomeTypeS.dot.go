// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// SomeTypeSChan represents a
// bidirectional
// channel
type SomeTypeSChan interface {
	SomeTypeSROnlyChan // aka "<-chan" - receive only
	SomeTypeSSOnlyChan // aka "chan<-" - send only
}

// SomeTypeSROnlyChan represents a
// receive-only
// channel
type SomeTypeSROnlyChan interface {
	RequestSomeTypeS() (dat []SomeType)        // the receive function - aka "MySomeTypeS := <-MySomeTypeSROnlyChan"
	TrySomeTypeS() (dat []SomeType, open bool) // the multi-valued comma-ok receive function - aka "MySomeTypeS, ok := <-MySomeTypeSROnlyChan"
}

// SomeTypeSSOnlyChan represents a
// send-only
// channel
type SomeTypeSSOnlyChan interface {
	ProvideSomeTypeS(dat []SomeType) // the send function - aka "MyKind <- some SomeTypeS"
}
