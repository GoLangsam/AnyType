// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Int8Chan represents a
// bidirectional
// channel
type Int8Chan interface {
	Int8ROnlyChan // aka "<-chan" - receive only
	Int8SOnlyChan // aka "chan<-" - send only
}

// Int8ROnlyChan represents a
// receive-only
// channel
type Int8ROnlyChan interface {
	RequestInt8() (dat int8)        // the receive function - aka "MyInt8 := <-MyInt8ROnlyChan"
	TryInt8() (dat int8, open bool) // the multi-valued comma-ok receive function - aka "MyInt8, ok := <-MyInt8ROnlyChan"
}

// Int8SOnlyChan represents a
// send-only
// channel
type Int8SOnlyChan interface {
	ProvideInt8(dat int8) // the send function - aka "MyKind <- some Int8"
}
