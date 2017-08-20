// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Complex128Chan represents a
// bidirectional
// channel
type Complex128Chan interface {
	Complex128ROnlyChan // aka "<-chan" - receive only
	Complex128SOnlyChan // aka "chan<-" - send only
}

// Complex128ROnlyChan represents a
// receive-only
// channel
type Complex128ROnlyChan interface {
	RequestComplex128() (dat complex128)        // the receive function - aka "MyComplex128 := <-MyComplex128ROnlyChan"
	TryComplex128() (dat complex128, open bool) // the multi-valued comma-ok receive function - aka "MyComplex128, ok := <-MyComplex128ROnlyChan"
}

// Complex128SOnlyChan represents a
// send-only
// channel
type Complex128SOnlyChan interface {
	ProvideComplex128(dat complex128) // the send function - aka "MyKind <- some Complex128"
}
