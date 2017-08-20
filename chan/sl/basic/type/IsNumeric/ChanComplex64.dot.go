// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Complex64Chan represents a
// bidirectional
// channel
type Complex64Chan interface {
	Complex64ROnlyChan // aka "<-chan" - receive only
	Complex64SOnlyChan // aka "chan<-" - send only
}

// Complex64ROnlyChan represents a
// receive-only
// channel
type Complex64ROnlyChan interface {
	RequestComplex64() (dat complex64)        // the receive function - aka "MyComplex64 := <-MyComplex64ROnlyChan"
	TryComplex64() (dat complex64, open bool) // the multi-valued comma-ok receive function - aka "MyComplex64, ok := <-MyComplex64ROnlyChan"
}

// Complex64SOnlyChan represents a
// send-only
// channel
type Complex64SOnlyChan interface {
	ProvideComplex64(dat complex64) // the send function - aka "MyKind <- some Complex64"
}
