// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsComplex

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type Complex64Chan interface { // bidirectional channel
	Complex64ROnlyChan // aka "<-chan" - receive only
	Complex64SOnlyChan // aka "chan<-" - send only
}

type Complex64ROnlyChan interface { // receive-only channel
	RequestComplex64() (dat complex64)        // the receive function - aka "some-new-Complex64-var := <-MyKind"
	TryComplex64() (dat complex64, open bool) // the multi-valued comma-ok receive function - aka "some-new-Complex64-var, ok := <-MyKind"
}

type Complex64SOnlyChan interface { // send-only channel
	ProvideComplex64(dat complex64) // the send function - aka "MyKind <- some Complex64"
}
