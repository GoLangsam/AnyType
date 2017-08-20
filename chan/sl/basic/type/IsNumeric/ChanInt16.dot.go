// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Int16Chan represents a
// bidirectional
// channel
type Int16Chan interface {
	Int16ROnlyChan // aka "<-chan" - receive only
	Int16SOnlyChan // aka "chan<-" - send only
}

// Int16ROnlyChan represents a
// receive-only
// channel
type Int16ROnlyChan interface {
	RequestInt16() (dat int16)        // the receive function - aka "MyInt16 := <-MyInt16ROnlyChan"
	TryInt16() (dat int16, open bool) // the multi-valued comma-ok receive function - aka "MyInt16, ok := <-MyInt16ROnlyChan"
}

// Int16SOnlyChan represents a
// send-only
// channel
type Int16SOnlyChan interface {
	ProvideInt16(dat int16) // the send function - aka "MyKind <- some Int16"
}
