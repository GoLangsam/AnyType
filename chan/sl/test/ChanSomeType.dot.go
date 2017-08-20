// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// SomeTypeChan represents a
// bidirectional
// channel
type SomeTypeChan interface {
	SomeTypeROnlyChan // aka "<-chan" - receive only
	SomeTypeSOnlyChan // aka "chan<-" - send only
}

// SomeTypeROnlyChan represents a
// receive-only
// channel
type SomeTypeROnlyChan interface {
	RequestSomeType() (dat SomeType)        // the receive function - aka "MySomeType := <-MySomeTypeROnlyChan"
	TrySomeType() (dat SomeType, open bool) // the multi-valued comma-ok receive function - aka "MySomeType, ok := <-MySomeTypeROnlyChan"
}

// SomeTypeSOnlyChan represents a
// send-only
// channel
type SomeTypeSOnlyChan interface {
	ProvideSomeType(dat SomeType) // the send function - aka "MyKind <- some SomeType"
}
