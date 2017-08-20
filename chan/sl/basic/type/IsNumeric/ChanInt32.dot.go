// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Int32Chan represents a
// bidirectional
// channel
type Int32Chan interface {
	Int32ROnlyChan // aka "<-chan" - receive only
	Int32SOnlyChan // aka "chan<-" - send only
}

// Int32ROnlyChan represents a
// receive-only
// channel
type Int32ROnlyChan interface {
	RequestInt32() (dat int32)        // the receive function - aka "MyInt32 := <-MyInt32ROnlyChan"
	TryInt32() (dat int32, open bool) // the multi-valued comma-ok receive function - aka "MyInt32, ok := <-MyInt32ROnlyChan"
}

// Int32SOnlyChan represents a
// send-only
// channel
type Int32SOnlyChan interface {
	ProvideInt32(dat int32) // the send function - aka "MyKind <- some Int32"
}
