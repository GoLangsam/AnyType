// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type Int64Chan interface { // bidirectional channel
	Int64ROnlyChan // aka "<-chan" - receive only
	Int64SOnlyChan // aka "chan<-" - send only
}

type Int64ROnlyChan interface { // receive-only channel
	RequestInt64() (dat int64)        // the receive function - aka "some-new-Int64-var := <-MyKind"
	TryInt64() (dat int64, open bool) // the multi-valued comma-ok receive function - aka "some-new-Int64-var, ok := <-MyKind"
}

type Int64SOnlyChan interface { // send-only channel
	ProvideInt64(dat int64) // the send function - aka "MyKind <- some Int64"
}
