// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Int64Chan represents a
// bidirectional
// channel
type Int64Chan interface {
	Int64ROnlyChan // aka "<-chan" - receive only
	Int64SOnlyChan // aka "chan<-" - send only
}

// Int64ROnlyChan represents a
// receive-only
// channel
type Int64ROnlyChan interface {
	RequestInt64() (dat int64)        // the receive function - aka "MyInt64 := <-MyInt64ROnlyChan"
	TryInt64() (dat int64, open bool) // the multi-valued comma-ok receive function - aka "MyInt64, ok := <-MyInt64ROnlyChan"
}

// Int64SOnlyChan represents a
// send-only
// channel
type Int64SOnlyChan interface {
	ProvideInt64(dat int64) // the send function - aka "MyKind <- some Int64"
}
