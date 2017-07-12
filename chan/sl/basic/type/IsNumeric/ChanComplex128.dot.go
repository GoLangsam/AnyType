// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type Complex128Chan interface { // bidirectional channel
	Complex128ROnlyChan // aka "<-chan" - receive only
	Complex128SOnlyChan // aka "chan<-" - send only
}

type Complex128ROnlyChan interface { // receive-only channel
	RequestComplex128() (dat complex128)        // the receive function - aka "some-new-Complex128-var := <-MyKind"
	TryComplex128() (dat complex128, open bool) // the multi-valued comma-ok receive function - aka "some-new-Complex128-var, ok := <-MyKind"
}

type Complex128SOnlyChan interface { // send-only channel
	ProvideComplex128(dat complex128) // the send function - aka "MyKind <- some Complex128"
}
